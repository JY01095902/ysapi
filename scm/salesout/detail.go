package salesout

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type DetailRequest struct {
	AppKey    string
	AppSecret string
	Id        string
}

type DetailResponse struct {
	Code    string    `json:"code"`
	Message string    `json:"message"`
	Data    DetailDto `json:"data"`
}

func Get(req DetailRequest) (DetailResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Get(request.URLRoot+"/yonbip/scm/salesout/detail", map[string]string{
		"id": req.Id,
	})
	if err != nil {
		return DetailResponse{}, err
	}

	resp, err := vals.GetResult(DetailResponse{})
	if err != nil {
		return DetailResponse{}, err
	}

	res, ok := resp.(*DetailResponse)
	if !ok {
		return DetailResponse{}, fmt.Errorf("%w error: response is not type of DetailResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	return *res, nil
}

type DetailValuesResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

// response 里用 values
func Detail(req DetailRequest) (DetailValuesResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Get(request.URLRoot+"/yonbip/scm/salesout/detail", map[string]string{
		"id": req.Id,
	})
	if err != nil {
		return DetailValuesResponse{}, err
	}

	resp, err := vals.GetResult(DetailValuesResponse{})
	if err != nil {
		return DetailValuesResponse{}, err
	}

	res, ok := resp.(*DetailValuesResponse)
	if !ok {
		return DetailValuesResponse{}, fmt.Errorf("%w error: response is not type of DetailValuesResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	return *res, nil
}
