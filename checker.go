package ysapi

import "fmt"

type ListResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func checkAPIResponse(val Values) error {
	result, err := val.GetResult(ListResponse{})
	if err != nil {
		return fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, err.Error())
	}

	resp, ok := result.(*ListResponse)
	if !ok {
		return fmt.Errorf("%w: result is not ListResponse", ErrYonSuiteAPIBizError)
	}

	if resp.Code != "200" {
		return fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, resp.Message)
	}

	return nil
}
