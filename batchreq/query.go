package batchreq

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type result struct {
	num  int
	data interface{}
	err  error
}

func bundelErrors(errs ...error) error {
	msg := ""
	for _, err := range errs {
		if err != nil {
			msg += ";" + err.Error()
		}
	}

	if msg != "" {
		return errors.New(strings.TrimPrefix(msg, ";"))
	}

	return nil
}

func newNumsChannel(count int) chan int {
	ch := make(chan int, count)
	for i := 1; i <= count; i++ {
		ch <- i
	}

	return ch
}

func newResults(ctx context.Context, count int, results *[]interface{}, errs *[]error, nums chan<- int) chan<- result {
	ch := make(chan result, count)

	go func() {
		defer close(nums)

		resmap := make(map[int]struct{}, count) // 用于判断查询是否全部完成
		for {
			select {
			case res, ok := <-ch:
				if !ok {
					return
				}

				if res.err != nil {
					(*errs)[res.num-1] = res.err
					nums <- res.num
				} else {
					(*results)[res.num-1] = res.data
					(*errs)[res.num-1] = nil
					resmap[res.num] = struct{}{}

					if len(resmap) == count {
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return ch
}

func Query(count int, do func(index int) (interface{}, error), limiter *rate.Limiter, timeout time.Duration) ([]interface{}, error) {
	if count <= 0 {
		return []interface{}{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make([]interface{}, count)
	errs := make([]error, count)

	nums := newNumsChannel(count)
	reschann := newResults(ctx, count, &results, &errs, nums)
	var limiterErr error

	var wg sync.WaitGroup
ReadNum:
	for {
		select {
		case num, ok := <-nums:
			if !ok {
				break ReadNum
			}

			err := limiter.Wait(ctx)
			if err != nil {
				limiterErr = err
				break ReadNum
			}

			wg.Add(1)
			go func() {
				defer wg.Done()

				data, err := do(num)
				reschann <- result{
					num:  num,
					data: data,
					err:  err,
				}
			}()
		case <-ctx.Done():
			break ReadNum
		}
	}

	wg.Wait()
	defer close(reschann)

	return results, bundelErrors(append(errs, ctx.Err(), limiterErr)...)
}

// func Query111(count int, do func(index int) (interface{}, error), limiter *rate.Limiter, timeout time.Duration) ([]interface{}, error) {
// 	if count <= 0 {
// 		return []interface{}{}, nil
// 	}

// 	start := time.Now()
// 	type result struct {
// 		num  int
// 		data interface{}
// 		err  error
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 	defer cancel()

// 	numChan := make(chan int, count)
// 	for i := 1; i <= count; i++ {
// 		numChan <- i
// 	}

// 	results := make([]interface{}, count)
// 	resultChan := make(chan result, count)
// 	errs := make([]error, count)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go func(ctx context.Context) {
// 		defer wg.Done()

// 		resmap := make(map[int]struct{}, count) // 用于判断查询是否全部完成
// 		for {
// 			select {
// 			case res, ok := <-resultChan:
// 				if !ok {
// 					return
// 				}

// 				if res.err != nil {
// 					errs[res.num-1] = res.err
// 					numChan <- res.num
// 				} else {
// 					results[res.num-1] = res.data
// 					errs[res.num-1] = nil
// 					resmap[res.num] = struct{}{}

// 					log.Printf("num: %d is done.", res.num)
// 					if len(resmap) == count {
// 						close(numChan)
// 						close(resultChan)
// 						return
// 					}
// 				}
// 			case <-ctx.Done():
// 				return
// 			}
// 		}
// 	}(ctx)

// 	var limiterErr error
// 	go func(ctx context.Context) {
// 		defer wg.Done()

// 		for {
// 			select {
// 			case num, ok := <-numChan:
// 				if !ok {
// 					return
// 				}
// 				err := limiter.Wait(context.Background())
// 				if err != nil {
// 					limiterErr = err
// 					return
// 				}

// 				go func() {
// 					data, err := do(num)
// 					select {
// 					case resultChan <- result{
// 						num:  num,
// 						data: data,
// 						err:  err,
// 					}:
// 					case <-ctx.Done():
// 						return
// 					}
// 				}()
// 			case <-ctx.Done():
// 				return
// 			}
// 		}
// 	}(ctx)

// 	wg.Wait()

// 	errMsg := ""
// 	if ctx.Err() != nil {
// 		errMsg += ";" + ctx.Err().Error()
// 	}

// 	if limiterErr != nil {
// 		errMsg += ";" + limiterErr.Error()
// 	}

// 	for _, err := range errs {
// 		if err != nil {
// 			errMsg += ";" + err.Error()
// 		}
// 	}

// 	log.Printf("err msg: %v", errMsg)
// 	// spew.Dump("results: ", results)
// 	log.Printf("duration: %s", time.Since(start))

// 	if errMsg != "" {
// 		return results, errors.New(strings.TrimPrefix(errMsg, ";"))
// 	}

// 	return results, nil
// }
