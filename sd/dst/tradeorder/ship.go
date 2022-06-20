package tradeorder

import (
	"fmt"

	"github.com/jy01095902/ysapi/request"
)

type ShipRequest struct {
	AppKey       string
	AppSecret    string
	PartParam    request.Values
	Path         int
	Ids          string // 多个的话逗号间隔
	ExternalData struct {
		ShipDetails  []request.Values
		Option       request.Values
		Modify       []request.Values
		ExpressLists []request.Values
	}
}

func (req ShipRequest) ToValues() request.Values {
	values := request.Values{
		"partParam": req.PartParam,
		"path":      req.Path,
		"ids":       req.Ids,
	}

	extData := request.Values{}
	if len(req.ExternalData.Option) > 0 {
		extData.Set("option", req.ExternalData.Option)
	}

	if len(req.ExternalData.Modify) > 0 {
		extData.Set("Modify", req.ExternalData.Modify)
	}

	extData.Set("shipdetails", req.ExternalData.ShipDetails)
	extData.Set("expresslists", req.ExternalData.ExpressLists)

	values.Set("externalData", extData)

	return values
}

type ShipResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func Ship(req ShipRequest) (ShipResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)
	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/tradeorder/ship", req.ToValues())
	if err != nil {
		return ShipResponse{}, err
	}

	res, err := vals.GetResult(ShipResponse{})
	if err != nil {
		return ShipResponse{}, err
	}

	resp, ok := res.(*ShipResponse)
	if !ok {
		return ShipResponse{}, fmt.Errorf("%w error: response is not type of ShipResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
