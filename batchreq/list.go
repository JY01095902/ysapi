package batchreq

import (
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

	getPageCount := func() int {
		if resp, err := listByPage(1); err == nil {
			log.Printf("total: %v", resp.Total())
			log.Printf("page count: %v", resp.PageCount())

			return resp.PageCount()
		}

		return 0
	}

	result, err := Query(getPageCount(), fetchListByPage, limiter, timeout)
	if err != nil {
		log.Printf("list all failed, error: %v", err)

		return nil, err
	}

	return result, nil
}
