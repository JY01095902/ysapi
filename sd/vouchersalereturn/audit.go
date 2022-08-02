package vouchersalereturn

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type AuditRequest struct {
	AppKey    string
	AppSecret string
	Data      request.Values
}

func (req AuditRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type AuditResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Audit(req AuditRequest) (AuditResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.NewURLRoot+"/sd/vouchersalereturn/audit", req.ToValues())
	if err != nil {
		return AuditResponse{}, err
	}

	res, err := vals.GetResult(AuditResponse{})
	if err != nil {
		return AuditResponse{}, err
	}

	resp, ok := res.(*AuditResponse)
	if !ok {
		return AuditResponse{}, fmt.Errorf("%w error: response is not type of AuditResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
