package refundorder

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gitlab.libratone.com/internet/ysapi.git/request"
)

type UnauditRequest struct {
	AppKey    string
	AppSecret string
	Ids       string
	PartParam request.Values
}

func (req UnauditRequest) ToValues() request.Values {
	values := request.Values{
		"ids":       req.Ids,
		"partParam": req.PartParam,
	}

	return values
}

/*
	{
	    "code": "200",
	    "message": "[{\"actionName\":\"退换货单弃审\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":true,\"sucIdAndPubts\":{\"1588813225906405382\":\"2022-11-14 09:52:37\"},\"successCount\":\"1\"},{\"actionName\":\"取消奇门退换货单据\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1588813225906405382\":\"2022-11-14 09:52:37\"},\"successCount\":\"0\"},{\"actionName\":\"取消顺丰WMS入库单\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1588813225906405382\":\"2022-11-14 09:52:37\"},\"successCount\":\"0\"},{\"actionName\":\"提交存量\",\"code\":\"1\",\"externalMap\":{},\"failCount\":\"0\",\"isExcuteAction\":true,\"isShowMsg\":false,\"sucIdAndPubts\":{\"1588813225906405382\":\"2022-11-14 09:52:37\"},\"successCount\":\"0\"}]",
	    "data": null
	}
*/
type UnauditResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"message"`
	Data    []request.Values `json:"data"`
}

func (resp UnauditResponse) IsSuccessed(id string) (bool, string) {
	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return false, err.Error()
	}
	for _, action := range actions {
		_, extid := action.SucIdAndPubts[id]
		if action.ActionName == "退换货单弃审" {
			if action.SuccessCount == "1" && extid {
				return true, ""
			}

			if action.FailCount == "1" {
				return false, action.ExceptionMsg
			}
		}
	}

	// 没有订单弃审的事件，把data的内容都返回
	return false, resp.Message
}

func (resp UnauditResponse) Timestamp(id string) string {
	if ok, _ := resp.IsSuccessed(id); !ok {
		return ""
	}

	var actions []action
	err := json.Unmarshal([]byte(resp.Message), &actions)
	if err != nil {
		return ""
	}
	for _, action := range actions {
		if action.ActionName == "退换货单弃审" {
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

func Unaudit(req UnauditRequest) (UnauditResponse, error) {
	apiReq := request.New(req.AppKey, req.AppSecret)

	vals, err := apiReq.Post(request.URLRoot+"/yonbip/sd/dst/refundorder/unaudit", req.ToValues())
	if err != nil {
		return UnauditResponse{}, err
	}

	res, err := vals.GetResult(UnauditResponse{})
	if err != nil {
		return UnauditResponse{}, err
	}

	resp, ok := res.(*UnauditResponse)
	if !ok {
		return UnauditResponse{}, fmt.Errorf("%w error: response is not type of UnauditResponse", request.ErrCallYonSuiteAPIFailed)
	}

	if resp.Code != "200" {
		return *resp, fmt.Errorf("%w error: %s", request.ErrCallYonSuiteAPIFailed, resp.Message)
	}

	return *resp, nil
}
