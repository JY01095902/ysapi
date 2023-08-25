package storeout

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type DetailDto struct {
	Id                   int64  `json:"id"`
	OrderNumber          string `json:"code"`
	SourceOrderNumber    string `json:"srcBillNO"`
	HeadItemDefine1      string `json:"headItem!define1"`
	HeadItemDefine2      string `json:"headItem!define2"`
	ReceiverMobileNumber string `json:"receivemobile"`
	OutWarehouseName     string `json:"outwarehouse_name"`
	Items                []struct {
		ProductSKUCode string  `json:"productsku_cCode"`
		Qty            float64 `json:"qty"`
		SerialNumbers  []struct {
			SerialNumber string `json:"sn"`
		} `json:"storeOutDetailSNs"`
	} `json:"details"`
}

type DetailRequest struct {
	AppKey    string
	AppSecret string
	Id        string
}

type DetailResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    DetailDto `json:"data"`
}

func Get(req DetailRequest) (DetailResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Get(request.URLRoot+"/yonbip/scm/storeout/detail", map[string]string{
		"id": req.Id,
	})
	if err != nil {
		return DetailResponse{}, err
	}

	resp, err := vals.GetResult(DetailResponse{})
	if err != nil {
		return DetailResponse{}, err
	}

	res, ok := resp.(*DetailResponse)
	if !ok {
		return DetailResponse{}, fmt.Errorf("%w error: response is not type of DetailResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	return *res, nil
}
