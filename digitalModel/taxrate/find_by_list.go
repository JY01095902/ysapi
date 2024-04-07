package taxrate

import (
	"fmt"
	"math"

	"github.com/jy01095902/ysapi/request"
)

type FindByListRequest struct {
	AppKey    string
	AppSecret string
	Params    request.Values
}

func (req FindByListRequest) ToValues() request.Values {
	values := request.Values{}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	return values
}

type FindByListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex  int              `json:"pageIndex"`
		PageSize   int              `json:"pageSize"`
		TotalCount int              `json:"recordCount"`
		Items      []request.Values `json:"recordList"`
	} `json:"data"`
}

func (resp FindByListResponse) Total() int {
	return resp.Data.TotalCount
}

func (resp FindByListResponse) PageCount() int {
	cnt := float64(resp.Data.TotalCount) / float64(resp.Data.PageSize)

	return int(math.Ceil(cnt))
}

func FindByList(req FindByListRequest) (FindByListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/digitalModel/taxrate/findByList", req.ToValues())
	if err != nil {
		return FindByListResponse{}, err
	}

	res, err := vals.GetResult(FindByListResponse{})
	if err != nil {
		return FindByListResponse{}, err
	}

	resp, ok := res.(*FindByListResponse)
	if !ok {
		return FindByListResponse{}, fmt.Errorf("%w error: response is not type of FindByListResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
