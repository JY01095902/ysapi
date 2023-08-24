package request

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type URLRootResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		GatewayURL string `json:"gatewayUrl"`
		TokenURL   string `json:"tokenUrl"`
	} `json:"data"`
}

func getURLRoot(tenantId string) (URLRootResponse, error) {
	cli := resty.New()
	resp, err := cli.R().
		SetQueryParam("tenantId", tenantId).
		SetResult(URLRootResponse{}).
		Get("https://apigateway.yonyoucloud.com/open-auth/dataCenter/getGatewayAddress")

	if err != nil {
		return URLRootResponse{}, err
	}

	if resp.StatusCode() != 200 {
		return URLRootResponse{}, fmt.Errorf("%w error: %s", ErrCallYonSuiteAPIFailed, resp.String())
	}

	result, ok := resp.Result().(*URLRootResponse)
	if !ok {
		return URLRootResponse{}, fmt.Errorf("%w: result is not url root response", ErrYonSuiteAPIBizError)
	}

	if result.Code != "00000" {
		return URLRootResponse{}, fmt.Errorf("%w: %s", ErrYonSuiteAPIBizError, result.Message)
	}

	return *result, nil
}
