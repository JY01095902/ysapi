package tradeorder

import (
	"encoding/json"
	"fmt"
	"strconv"

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

/*
{
    "code": "999",
    "message": " 未财审订单不允许发货!"
}

{
    "code": "200",
    "message": "[{\"actionName\":\"订单发货\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1576078312869986334\":\"2022-10-25 16:36:27\"},\"successCount\":\"1\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1576078312869986334\":\"2022-10-25 16:36:27\"},\"successCount\":\"1\"}]",
    "data": null
}
*/
type ShipResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func (resp ShipResponse) IsSuccessed(id string) (bool, string) {
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
		if action.ActionName == "订单发货" {
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

func (resp ShipResponse) Timestamp(id string) string {
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
		if action.ActionName == "订单发货" {
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

func Ship(req ShipRequest) (ShipResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	// fmt.Printf(" req.ToValues(): %s", req.ToValues().String())
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

func ShipAndUpload(req ShipRequest) (ShipResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/tradeorder/shipandupload", req.ToValues())
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

func Upload(req ShipRequest) (ShipResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/sd/dst/tradeorder/upload", req.ToValues())
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
