package batchreq

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/time/rate"
)

type ListResponse interface {
	Total() int
	PageCount() int
}

type ListByPage func(pageNumber int) (ListResponse, error)

func ListAll(listByPage ListByPage, limiter *rate.Limiter, timeout time.Duration) ([]interface{}, error) {
	fetchListByPage := func(num int) (interface{}, error) {
		log.Printf("page num: %v", num)

		return listByPage(num)
	}

	qres, err := Query(1, fetchListByPage, limiter, timeout)
	if err != nil {
		return nil, err
	}

	resp, ok := qres[0].(ListResponse)
	if !ok {
		return nil, fmt.Errorf("query result is not ListResponse")
	}

	result, err := Query(resp.PageCount(), fetchListByPage, limiter, timeout)
	if err != nil {
		log.Printf("list all failed, error: %v", err)

		return nil, err
	}

	return result, nil
}
