package batchreq

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"

	"github.com/jy01095902/ysapi/request"
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

func createResultChannel(ctx context.Context, count int) (<-chan int, chan<- result, <-chan struct{}, func() ([]interface{}, error)) {
	numch := make(chan int, count)
	for i := 1; i <= count; i++ {
		numch <- i
	}

	resultch := make(chan result, count)
	finishSignal := make(chan struct{})
	results := make([]interface{}, count)
	errs := make([]error, count)

	go func() {
		defer close(numch)
		defer close(finishSignal)

		resmap := make(map[int]struct{}, count) // 用于判断查询是否全部完成
		for {
			select {
			case res, ok := <-resultch:
				if !ok {
					return
				}

				if res.err != nil {
					errs[res.num-1] = res.err
					if errors.Is(res.err, request.ErrYonSuiteAPIBizError) {
						results[res.num-1] = res.data
						resmap[res.num] = struct{}{}

						if len(resmap) == count {
							finishSignal <- struct{}{}

							return
						}
					} else {
						numch <- res.num
					}
				} else {
					results[res.num-1] = res.data
					errs[res.num-1] = nil
					resmap[res.num] = struct{}{}

					if len(resmap) == count {
						finishSignal <- struct{}{}

						return
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return numch, resultch, finishSignal, func() ([]interface{}, error) {
		return results, bundelErrors(errs...)
	}
}

func Query(count int, do func(index int) (interface{}, error), limiter *rate.Limiter, timeout time.Duration) ([]interface{}, error) {
	if count <= 0 {
		return []interface{}{}, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	numch, resultch, finishSignal, getResults := createResultChannel(ctx, count)
	defer close(resultch)

	var limiterErr error
	var wg sync.WaitGroup
ReadNum:
	for {
		select {
		case num, ok := <-numch:
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
				resultch <- result{
					num:  num,
					data: data,
					err:  err,
				}
			}()

		case <-finishSignal:
			break ReadNum

		case <-ctx.Done():
			break ReadNum
		}
	}
	wg.Wait()

	results, err := getResults()

	return results, bundelErrors(err, ctx.Err(), limiterErr)
}
