package batchreq

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type BatchRequest struct {
	request       request.Request
	url           string
	body          request.Values
	pageNumField  string // 页码字段名
	pageSizeField string // 每页条数字段名
	totalField    string // 总条数字段名
	pageCnt       int    // 总页数字段名
	pageSize      int
	total         int
}

func New(appKey, appSecret, url string, body request.Values, pageNumField, pageSizeField, totalField string) BatchRequest {
	pageSize := 50
	body.Set(pageSizeField, pageSize)
	batreq := BatchRequest{
		request:       request.New(appKey, appSecret),
		url:           url,
		body:          body,
		pageNumField:  pageNumField,
		pageSizeField: pageSizeField,
		totalField:    totalField,
		pageSize:      pageSize,
	}

	return batreq
}

func (batreq BatchRequest) PageCount() int {
	return batreq.pageCnt
}

// request.Values 是每次请求的返回值（每页）
func (batreq *BatchRequest) Execute() (<-chan request.Values, <-chan error) {
Prepare:
	firstPage, err := batreq.prepare()
	if err != nil && errors.Is(err, request.ErrAPILimit) {
		log.Println("查询第 1 页时被限流了，等待1分钟后再查询......")
		time.Sleep(time.Minute)
		goto Prepare
	}
	log.Printf("prepare finished. page count: %d", batreq.pageCnt)
	pageNumChan := make(chan int, batreq.pageCnt)
	for i := 2; i <= batreq.pageCnt; i++ { // 第一页已经有了，从第二页开始获取
		pageNumChan <- i
	}
	log.Println("already put page numbers into chan.")

	chanSize := 10
	if batreq.pageCnt > 1 {
		chanSize = batreq.pageCnt
	}
	resChan := make(chan request.Values, chanSize)
	errChan := make(chan error, chanSize)
	resChan <- firstPage
	errChan <- err

	gcnt := 1 // 一次请求的goroutine数量
	var wg sync.WaitGroup
	wg.Add(gcnt)
	go func() {
		wg.Wait()

		close(resChan)
		close(errChan)
		close(pageNumChan)
	}()
	for i := 0; i < gcnt; i++ {
		go func() {
			defer wg.Done()

			for pageNum := range pageNumChan {
				res, err := batreq.fetchByPage(pageNum)
				if err != nil {
					if errors.Is(err, request.ErrAPILimit) {
						// 被限流了，需要等待1分钟后重新查询
						pageNumChan <- pageNum
						log.Printf("查询第 %d 页时被限流了，等待1分钟后再查询......", pageNum)
						time.Sleep(time.Minute)
					} else {
						errChan <- err
					}
				}

				resChan <- res
			}
		}()
	}

	return resChan, errChan
}

func (batreq *BatchRequest) prepare() (request.Values, error) {
	firstPage, total, err := batreq.fetchFirstPage()
	if err != nil {
		return request.Values{}, err
	}

	batreq.setPageCount(total)

	return firstPage, nil
}

func (batreq *BatchRequest) setPageCount(total int) {
	size := float64(batreq.pageSize)
	tot := float64(total)
	batreq.pageCnt = int(math.Ceil(tot / size))
	log.Printf("size: %v", size)
	log.Printf("tot: %v", tot)
	log.Printf("cnt: %v", math.Ceil(tot/size))
	log.Printf("pageCnt: %v", batreq.pageCnt)
	batreq.total = total
}

func (batreq BatchRequest) copyBody() request.Values {
	newBody := request.Values{}
	for k, v := range batreq.body {
		newBody.Set(k, v)
	}

	return newBody
}

func (batreq BatchRequest) fetchFirstPage() (request.Values, int, error) {
	body := batreq.copyBody()
	body.Set(batreq.pageNumField, 1)
	res, err := batreq.request.Post(batreq.url, body)
	if err != nil {
		return request.Values{}, 0, err
	}

	if mdata, ok := res.Get("data").(map[string]interface{}); ok {
		data := request.Values(mdata)
		switch total := data.Get(batreq.totalField).(type) {
		case json.Number:
			t, err := strconv.Atoi(total.String())
			return res, t, err
		case int:
			return res, total, nil
		}
	}

	return res, 0, errors.New("can not parse items total")
}

func (batreq BatchRequest) fetchByPage(pageNum int) (request.Values, error) {
	body := batreq.copyBody()
	body.Set(batreq.pageNumField, pageNum)
	res, err := batreq.request.Post(batreq.url, body)
	if err != nil {
		return request.Values{}, err
	}

	return res, nil
}
