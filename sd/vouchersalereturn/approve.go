package vouchersalereturn

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type ApproveRequest struct {
	AppKey    string
	AppSecret string
	Data      request.Values
}

func (req ApproveRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type ApproveResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Approve(req ApproveRequest) (ApproveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.NewURLRoot+"/yonbip/sd/vouchersalereturn/approve", req.ToValues())
	if err != nil {
		return ApproveResponse{}, err
	}

	res, err := vals.GetResult(ApproveResponse{})
	if err != nil {
		return ApproveResponse{}, err
	}

	resp, ok := res.(*ApproveResponse)
	if !ok {
		return ApproveResponse{}, fmt.Errorf("%w error: response is not type of ApproveResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
