package snflowdirection

import (
	"fmt"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type ListRequest struct {
	AppKey    string
	AppSecret string
	PageIndex int
	PageSize  int
	Params    request.Values
}

func (req ListRequest) ToValues() request.Values {
	values := request.Values{
		"pageIndex": req.PageIndex,
		"pageSize":  req.PageSize,
	}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	return values
}

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex      int              `json:"pageIndex"`
		PageSize       int              `json:"pageSize"`
		RecordCount    int              `json:"recordCount"`
		RecordList     []request.Values `json:"recordList"`
		PageCount      int              `json:"pageCount"`
		BeginPageIndex int              `json:"beginPageIndex"`
		EndPageIndex   int              `json:"endPageIndex"`
	} `json:"data"`
}

func (resp ListResponse) Total() int {
	return resp.Data.RecordCount
}

func (resp ListResponse) PageCount() int {
	return resp.Data.PageCount
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/scm/snflowdirection/list", req.ToValues())
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

	if len(resp.Data.RecordList) == 0 {
		return *resp, fmt.Errorf("error: %w sn flow", request.ErrNotFound)
	}

	return *resp, nil
}
