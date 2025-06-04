package unit

import (
	"fmt"
	"math"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type ListRequest struct {
	AppKey    string
	AppSecret string
	Params    request.Values
}

func (req ListRequest) ToValues() request.Values {
	values := request.Values{}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	return values
}

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex  int              `json:"pageIndex"`
		PageSize   int              `json:"pageSize"`
		TotalCount int              `json:"recordCount"`
		Items      []request.Values `json:"recordList"`
	} `json:"data"`
}

func (resp ListResponse) Total() int {
	return resp.Data.TotalCount
}

func (resp ListResponse) PageCount() int {
	cnt := float64(resp.Data.TotalCount) / float64(resp.Data.PageSize)

	return int(math.Ceil(cnt))
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/digitalModel/unit/list", req.ToValues())
	if err != nil {
		return ListResponse{}, err
	}

	res, err := vals.GetResult(ListResponse{})
	if err != nil {
		return ListResponse{}, err
	}

	resp, ok := res.(*ListResponse)
	if !ok {
		return ListResponse{}, fmt.Errorf("%w error: response is not type of ListResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
