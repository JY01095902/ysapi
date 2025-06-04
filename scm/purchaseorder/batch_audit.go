package purchaseorder

import (
	"fmt"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type BatchAuditRequest struct {
	AppKey    string
	AppSecret string
	Data      []request.Values
}

func (req BatchAuditRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type BatchAuditResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func BatchAudit(req BatchAuditRequest) (BatchAuditResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/scm/purchaseorder/batchaudit", req.ToValues())
	if err != nil {
		return BatchAuditResponse{}, err
	}

	res, err := vals.GetResult(BatchAuditResponse{})
	if err != nil {
		return BatchAuditResponse{}, err
	}

	resp, ok := res.(*BatchAuditResponse)
	if !ok {
		return BatchAuditResponse{}, fmt.Errorf("%w error: response is not type of BatchAuditResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
