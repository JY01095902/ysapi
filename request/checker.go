package request

import (
	"fmt"
)

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func checkResponse(val Values) error {
	result, err := val.GetResult(response{})
	if err != nil {
		return fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, err.Error())
	}

	resp, ok := result.(*response)
	if !ok {
		return fmt.Errorf("%w: result is not response", ErrYonSuiteAPIBizError)
	}

	if resp.Code == "200" {
		return nil
	}

	if resp.Code == "310046" {
		return fmt.Errorf("%w: %s", ErrAPILimit, resp.Message)
	}

	return fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, resp.Message)
}
