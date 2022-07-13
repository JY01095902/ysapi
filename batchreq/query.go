package batchreq

import (
	"context"
	"errors"
	"log"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func Query(count int, do func(index int) (interface{}, error), limiter *rate.Limiter, timeout time.Duration) ([]interface{}, error) {
	start := time.Now()
	type result struct {
		num  int
		data interface{}
		err  error
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	numChan := make(chan int, count)
	for i := 1; i <= count; i++ {
		numChan <- i
	}

	results := make([]interface{}, count)
	resultChan := make(chan result, count)
	errs := make([]error, count)
	var wg sync.WaitGroup
	wg.Add(2)
	go func(ctx context.Context) {
		defer wg.Done()

		resmap := make(map[int]struct{}, count) // 用于判断查询是否全部完成
		for {
			select {
			case res, ok := <-resultChan:
				if !ok {
					return
				}

				if res.err != nil {
					errs[res.num-1] = res.err
					numChan <- res.num
				} else {
					results[res.num-1] = res.data
					errs[res.num-1] = nil
					resmap[res.num] = struct{}{}

					log.Printf("num: %d is done.", res.num)
					if len(resmap) == count {
						close(numChan)
						close(resultChan)
						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	var limiterErr error
	go func(ctx context.Context) {
		defer wg.Done()

		for {
			select {
			case num, ok := <-numChan:
				if !ok {
					return
				}
				err := limiter.Wait(context.Background())
				if err != nil {
					limiterErr = err
					return
				}

				go func() {
					data, err := do(num)
					select {
					case resultChan <- result{
						num:  num,
						data: data,
						err:  err,
					}:
					case <-ctx.Done():
						return
					}
				}()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	wg.Wait()

	errMsg := ""
	if ctx.Err() != nil {
		errMsg += ";" + ctx.Err().Error()
	}

	if limiterErr != nil {
		errMsg += ";" + limiterErr.Error()
	}

	for _, err := range errs {
		if err != nil {
			errMsg += ";" + err.Error()
		}
	}

	log.Printf("err msg: %v", errMsg)
	// spew.Dump("results: ", results)
	log.Printf("duration: %s", time.Since(start))

	if errMsg != "" {
		return results, errors.New(strings.TrimPrefix(errMsg, ";"))
	}

	return results, nil
}
