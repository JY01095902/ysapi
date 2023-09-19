package merchant

import (
	"encoding/json"
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type NewUpdateRequest struct {
	AppKey    string
	AppSecret string
	Params    request.Values
}

func (req NewUpdateRequest) ToValues() request.Values {
	values := request.Values{}
	for k, v := range req.Params {
		values.Set(k, v)
	}

	return values
}

type NewUpdateResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id json.Number `json:"id"`
	} `json:"data"`
}

func NewUpdate(req NewUpdateRequest) (NewUpdateResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/yonbip/digitalModel/merchant/newupdate", req.ToValues())
	if err != nil {
		return NewUpdateResponse{}, err
	}

	resp, err := vals.GetResult(NewUpdateResponse{})
	if err != nil {
		return NewUpdateResponse{}, err
	}

	res, ok := resp.(*NewUpdateResponse)
	if !ok {
		return NewUpdateResponse{}, fmt.Errorf("%w error: response is not type of NewUpdateResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if res.Code != "200" {
		return *res, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, res.Message)
	}

	return *res, nil
}
