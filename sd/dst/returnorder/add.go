package returnorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jy01095902/ysapi/request"
)

type AddRequest struct {
	AppKey       string
	AppSecret    string
	PartParam    request.Values
	ExternalData struct {
		Head  []request.Values
		Added []request.Values
	}
}

func (req AddRequest) ToValues() request.Values {
	values := request.Values{
		"partParam": req.PartParam,
		"externalData": request.Values{
			"Head":  req.ExternalData.Head,
			"Added": req.ExternalData.Added,
		},
	}

	return values
}

/*
	{
	    "code": "200",
	    "message": "[{\"actionName\":\"订单挂起\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"1\"},{\"actionName\":\"取消顺丰WMS出库单\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"0\"},{\"actionName\":\"取消奇门单据\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509293890716303371\":1658905111000},\"successCount\":\"0\"},{\"actionName\":\"退换货单新增\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1509295411134726149\":1658905111000},\"successCount\":\"1\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1509295411134726149\":1658905111000},\"successCount\":\"1\"}]",
	    "data": null
	}
*/
type AddResponse struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Data    request.Values `json:"data"`
}

func (resp AddResponse) IsSuccessed() (bool, string) {
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
		// id 是创建之后才有的之前不知道
		// _, extid := action.SucIdAndPubts[id]
		if action.ActionName == "退换货单新增" {
			if action.SuccessCount == "1" {
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

func (resp AddResponse) Timestamp() string {
	if ok, _ := resp.IsSuccessed(); !ok {
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
		if action.ActionName == "退换货单新增" {
			// 默认第一个
			// id 是创建之后才有的之前不知道
			for _, v := range action.SucIdAndPubts {
				switch val := v.(type) {
				case int:
					return strconv.Itoa(val)
				case int64:
					return strconv.FormatInt(val, 10)
				case float64:
					return strconv.FormatInt(int64(val), 10)
				case string:
					return val
				default:
					return ""
				}
			}
		}
	}

	return ""
}

func Add(req AddRequest) (AddResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/returnorder/add", req.ToValues())
	if err != nil {
		return AddResponse{}, err
	}

	res, err := vals.GetResult(AddResponse{})
	if err != nil {
		return AddResponse{}, err
	}

	resp, ok := res.(*AddResponse)
	if !ok {
		return AddResponse{}, fmt.Errorf("%w error: response is not type of AddResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
