package admindept

import (
	"fmt"
	"log"

	"github.com/jy01095902/ysapi/request"
)

type ListRequest struct {
	AppKey       string
	AppSecret    string
	Params       request.Values
	ExternalData request.Values
}

func (req ListRequest) ToValues() request.Values {
	values := request.Values{
		"externalData": req.ExternalData,
	}

	for k, v := range req.Params {
		values.Set(k, v)
	}

	log.Printf("req: %v", values)
	return values
}

type ListResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    []request.Values `json:"data"`
}

func List(req ListRequest) (ListResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/digitalModel/admindept/tree", req.ToValues())
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

	if len(res.Data) == 0 {
		return *res, fmt.Errorf("%w error: not found department list", request.ErrYonSuiteAPIBizError)
	}

	return *res, nil
}
