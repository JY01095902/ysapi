package batchreq

import (
	"errors"
	"math"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
)

func TestQuery(t *testing.T) {
	type page struct {
		PageNumber int
		Items      []string
	}
	pageSize := 20
	itemCount := 205
	pageCount := int(math.Ceil(float64(itemCount) / float64(pageSize)))
	timeout := 1 * time.Second
	limiter := rate.NewLimiter(500, 1000)
	items := make([]string, itemCount)
	for i := 0; i < itemCount; i++ {
		items[i] = "DATA[" + strconv.Itoa(i) + "]"
	}

	// 正常成功
	do := func(pnum int) (interface{}, error) {
		start := (pnum - 1) * pageSize
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}

		return page{
			PageNumber: pnum,
			Items:      items[start:end],
		}, nil
	}

	result, err := Query(pageCount, do, limiter, timeout)
	assert.Nil(t, err)
	assert.Equal(t, pageCount, len(result))
	total := 0
	for _, inf := range result {
		if inf == nil {
			continue
		}
		p := inf.(page)
		total += len(p.Items)
	}
	assert.Equal(t, itemCount, total)

	// 一直返回错误，延时
	do = func(pnum int) (interface{}, error) {
		start := (pnum - 1) * pageSize
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}

		return page{
			PageNumber: pnum,
			Items:      items[start:end],
		}, errors.New("限流")
	}

	_, err = Query(pageCount, do, limiter, timeout)
	assert.NotNil(t, err)

	// 失败后重试成功
	var doCnt sync.Map
	do = func(pnum int) (interface{}, error) {
		start := (pnum - 1) * pageSize
		end := start + pageSize
		if end > len(items) {
			end = len(items)
		}
		data := items[start:end]
		if _, exist := doCnt.Load(pnum); exist {
			return page{
				PageNumber: pnum,
				Items:      data,
			}, nil
		}
		doCnt.Store(pnum, struct{}{})

		return page{
			PageNumber: pnum,
			Items:      data,
		}, errors.New("限流")
	}
	result, err = Query(pageCount, do, limiter, timeout)
	assert.Nil(t, err)
	assert.Equal(t, pageCount, len(result))
	total = 0
	for _, inf := range result {
		if inf == nil {
			continue
		}
		p := inf.(page)
		total += len(p.Items)
	}
	assert.Equal(t, itemCount, total)

	// 只执行一次
	result, err = Query(1, do, limiter, timeout)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
	total = 0
	for _, inf := range result {
		if inf == nil {
			continue
		}
		p := inf.(page)
		total += len(p.Items)
	}
	assert.Equal(t, pageSize, total)

	// 执行0次
	result, err = Query(0, do, limiter, timeout)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(result))
}
