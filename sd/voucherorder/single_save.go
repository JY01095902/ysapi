package voucherorder

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type SingleSaveRequest struct {
	AppKey    string
	AppSecret string
	Data      request.Values
}

func (req SingleSaveRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type SingleSaveResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func SingleSave(req SingleSaveRequest) (SingleSaveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/voucherorder/singleSave", req.ToValues())
	if err != nil {
		return SingleSaveResponse{}, err
	}

	res, err := vals.GetResult(SingleSaveResponse{})
	if err != nil {
		return SingleSaveResponse{}, err
	}

	resp, ok := res.(*SingleSaveResponse)
	if !ok {
		return SingleSaveResponse{}, fmt.Errorf("%w error: response is not type of SingleSaveResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
