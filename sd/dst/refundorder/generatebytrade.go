package refundorder

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type GenerateByTradeRequest struct {
	AppKey       string
	AppSecret    string
	Ids          string
	PartParam    request.Values
	ExternalData request.Values
}

func (req GenerateByTradeRequest) ToValues() request.Values {
	values := request.Values{
		"ids":          req.Ids,
		"partParam":    req.PartParam,
		"externalData": req.ExternalData,
	}

	return values
}

type GenerateByTradeResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    []request.Values `json:"data"`
}

func GenerateByTrade(req GenerateByTradeRequest) (GenerateByTradeResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/refundorder/generatebytrade", req.ToValues())
	if err != nil {
		return GenerateByTradeResponse{}, err
	}

	res, err := vals.GetResult(GenerateByTradeResponse{})
	if err != nil {
		return GenerateByTradeResponse{}, err
	}

	resp, ok := res.(*GenerateByTradeResponse)
	if !ok {
		return GenerateByTradeResponse{}, fmt.Errorf("%w error: response is not type of GenerateByTradeResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
