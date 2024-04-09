package mergesourcedata

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type SaveRequest struct {
	AppKey    string
	AppSecret string
	Data      request.Values
}

func (req SaveRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type SaveResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Save(req SaveRequest) (SaveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/scm/purinrecord/mergeSourceData/save", req.ToValues())
	if err != nil {
		return SaveResponse{}, err
	}

	res, err := vals.GetResult(SaveResponse{})
	if err != nil {
		return SaveResponse{}, err
	}

	resp, ok := res.(*SaveResponse)
	if !ok {
		return SaveResponse{}, fmt.Errorf("%w error: response is not type of SaveResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
