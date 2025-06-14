package storearchives

import (
	"fmt"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type QueryRequest struct {
	AppKey    string
	AppSecret string
	LikeValue string
	PartParam request.Values
}

func (req QueryRequest) ToValues() request.Values {
	values := request.Values{
		"partParam": req.PartParam,
		"likeValue": req.LikeValue,
	}

	return values
}

type QueryResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		PageIndex  int              `json:"pageIndex"`
		PageSize   int              `json:"pageSize"`
		TotalCount int              `json:"totalCount"`
		Infos      []request.Values `json:"info"`
	} `json:"data"`
}

func Query(req QueryRequest) (QueryResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/storearchives/query", req.ToValues())
	if err != nil {
		return QueryResponse{}, err
	}

	res, err := vals.GetResult(QueryResponse{})
	if err != nil {
		return QueryResponse{}, err
	}

	resp, ok := res.(*QueryResponse)
	if !ok {
		return QueryResponse{}, fmt.Errorf("%w error: response is not type of QueryResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
