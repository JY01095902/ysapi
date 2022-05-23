package salesout

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
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

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex      int             `json:"pageIndex"`
		PageSize       int             `json:"pageSize"`
		RecordCount    int             `json:"recordCount"`
		RecordList     []SalesOutOrder `json:"recordList"`
		PageCount      int             `json:"pageCount"`
		BeginPageIndex int             `json:"beginPageIndex"`
		EndPageIndex   int             `json:"endPageIndex"`
		PubTs          string          `json:"pubts"`
	} `json:"data"`
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.SalesOutListURL, req.ToValues())
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
		return *res, fmt.Errorf("%w error: not found sn state", request.ErrYonSuiteAPIBizError)
	}

	return *res, nil
}
