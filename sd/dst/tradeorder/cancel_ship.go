package tradeorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type CancelShipRequest struct {
	AppKey    string
	AppSecret string
	Ids       string
	PartParam request.Values
}

func (req CancelShipRequest) ToValues() request.Values {
	values := request.Values{
		"ids":       req.Ids,
		"partParam": req.PartParam,
	}

	return values
}

/*
	{
	    "code": "200",
	    "message": "[{\"actionName\":\"取消发货\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1632457936260825106\":\"2023-01-09 16:18:40\"},\"successCount\":\"1\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1632457936260825106\":\"2023-01-09 16:18:40\"},\"successCount\":\"1\"}]",
	    "data": null
	}
*/
type CancelShipResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func (resp CancelShipResponse) IsSuccessed(id string) (bool, string) {
	type action struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return false, err.Error()
	}
	for _, action := range actions {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "取消发货" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单发货的事件，把data的内容都返回
	return false, resp.Message
}

func (resp CancelShipResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	type action struct {
		ExceptionMsg   string         `json:"exceptionMsg"`
		Code           string         `json:"code"`
		IsShowMsg      bool           `json:"isShowMsg"`
		FailCount      string         `json:"failCount"`
		SucIdAndPubts  request.Values `json:"sucIdAndPubts"`
		SuccessCount   string         `json:"successCount"`
		IsExcuteAction bool           `json:"isExcuteAction"`
		ActionName     string         `json:"actionName"`
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return ""
	}
	for _, action := range actions {
		if action.ActionName == "取消发货" {
			switch val := action.SucIdAndPubts[id].(type) {
			case int:
				return strconv.Itoa(val)
			case int64:
				return strconv.FormatInt(val, 10)
			case float64:
				return strconv.FormatInt(int64(val), 10)
			case string:
				return val
			case json.Number:
				return val.String()
			default:
				return ""
			}
		}
	}

	return ""
}

func CancelShip(req CancelShipRequest) (CancelShipResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/tradeorder/cancelship", req.ToValues())
	if err != nil {
		return CancelShipResponse{}, err
	}

	res, err := vals.GetResult(CancelShipResponse{})
	if err != nil {
		return CancelShipResponse{}, err
	}

	resp, ok := res.(*CancelShipResponse)
	if !ok {
		return CancelShipResponse{}, fmt.Errorf("%w error: response is not type of CancelShipResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
