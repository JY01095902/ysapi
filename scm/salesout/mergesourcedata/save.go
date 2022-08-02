package mergesourcedata

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type Detail struct {
	SourceiId    string `json:"sourceid"`
	Sourceautoid string `json:"sourceautoid"`
	MakeRuleCode string `json:"makeRuleCode"`
	Status       string `json:"_status"`
}

type Order struct {
	MergeSourceData bool     `json:"mergeSourceData"`
	VouchDate       string   `json:"vouchdate"`
	Warehouse       string   `json:"warehouse"`
	Code            string   `json:"code"`
	BusType         string   `json:"bustype"`
	Details         []Detail `json:"details"`
	Status          string   `json:"_status"`
}

type SaveRequest struct {
	AppKey    string
	AppSecret string
	Data      Order
}

func (req SaveRequest) ToValues() request.Values {
	values := request.Values{
		"data": req.Data,
	}

	return values
}

type SaveResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Count        int              `json:"count"`
		SuccessCount int              `json:"sucessCount"`
		FailCount    int              `json:"failCount"`
		Messages     []string         `json:"messages"`
		Infos        []request.Values `json:"infos"`
	} `json:"data"`
}

func Save(req SaveRequest) (SaveResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.NewURLRoot+"/scm/salesout/mergeSourceData/save", req.ToValues())
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
