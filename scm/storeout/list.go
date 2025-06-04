package storeout

import (
	"errors"
	"fmt"
	"log"
	"sync"

	"gitlab.libratone.com/internet/ysapi.git/batchreq"
	"gitlab.libratone.com/internet/ysapi.git/request"
)

type SimpleVO struct {
	Field  string
	Op     string
	Value1 string
	Value2 string
}

type ListRequest struct {
	AppKey    string
	AppSecret string
	PageIndex int
	PageSize  int
	Params    request.Values
	SimpleVOs []SimpleVO
}

func (req ListRequest) ToValues() request.Values {
	values := request.Values{
		"pageIndex": req.PageIndex,
		"pageSize":  req.PageSize,
	}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	if len(req.SimpleVOs) > 0 {
		ovs := []request.Values{}
		for _, ov := range req.SimpleVOs {
			ovs = append(ovs, request.Values{
				"field":  ov.Field,
				"op":     ov.Op,
				"value1": ov.Value1,
				"value2": ov.Value2,
			})
		}

		values.Set("simpleVOs", ovs)
	}

	return values
}

type ListDto struct {
	Id                int64  `json:"id"`
	OrderNumber       string `json:"code"`
	SourceOrderNumber string `json:"srcbillno"`
}

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex      int       `json:"pageIndex"`
		PageSize       int       `json:"pageSize"`
		RecordCount    int       `json:"recordCount"`
		RecordList     []ListDto `json:"recordList"`
		PageCount      int       `json:"pageCount"`
		BeginPageIndex int       `json:"beginPageIndex"`
		EndPageIndex   int       `json:"endPageIndex"`
		PubTs          string    `json:"pubts"`
	} `json:"data"`
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/scm/storeout/list", req.ToValues())
	if err != nil {
		return ListResponse{}, err
	}

	resp, err := vals.GetResult(ListResponse{})
	if err != nil {
		return ListResponse{}, err
	}

	res, ok := resp.(*ListResponse)
	if !ok {
		return ListResponse{}, fmt.Errorf("%w error: response is not type of ListResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	if len(res.Data.RecordList) == 0 {
		return *res, fmt.Errorf("%w error: not found store-out list", request.ErrYonSuiteAPIBizError)
	}

	return *res, nil
}

func ListAll(req ListRequest) ([]ListResponse, error) {
	var list []ListResponse
	url := request.URLRoot + "/yonbip/scm/storeout/list"
	batreq := batchreq.New(req.AppKey, req.AppSecret, url, req.ToValues(), "pageIndex", "pageSize", "recordCount")
	// if batreq.ItemsTotal() == 0 {
	// 	return nil, errors.New("items total is 0")
	// }
	resChan, errChan := batreq.Execute()
	var err error
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		errstr := ""
		for e := range errChan {
			if e != nil {
				errstr += ";" + e.Error()
			}
		}
		if len(errstr) > 0 {
			err = errors.New(errstr[1:])
		}
	}()

	for res := range resChan {
		if list == nil {
			log.Printf("batreq.PageCount(): %d", batreq.PageCount())
			list = make([]ListResponse, batreq.PageCount())
		}

		resp, err := res.GetResult(ListResponse{})
		if err != nil {
			log.Printf("parse values to ListResponse failed, error: %s", err.Error())
		}

		lstResp, ok := resp.(*ListResponse)
		if !ok {
			log.Println("values is not type of ListResponse")
		}

		if lstResp.Data.PageIndex > 0 {
			list[lstResp.Data.PageIndex-1] = *lstResp
		}
	}

	wg.Wait()
	return list, err
}
