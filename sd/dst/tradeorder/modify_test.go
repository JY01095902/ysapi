package tradeorder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifyResponseIsSuccessed(t *testing.T) {
	resp := ModifyResponse{
		Code:    "",
		Message: "",
		Data:    map[string]interface{}{},
	}

	msgOK := `[{"actionName":"订单修改","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":true,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:53"},"successCount":"1"},{"actionName":"自动匹配","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"重量重算","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"表头汇总表体仓库","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"提交存量","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"}]`
	resp.Message = msgOK

	isSuccess, msg := resp.IsSuccessed()
	assert.Empty(t, msg)
	assert.True(t, isSuccess)

	msgErr := `[{"actionName":"订单修改","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":true,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:04:34"},"successCount":"1"},{"actionName":"自动匹配","code":"1","externalMap":{},"failCount":"1","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{},"successCount":"0"}]`
	resp.Message = msgErr

	isSuccess, msg = resp.IsSuccessed()
	assert.NotEmpty(t, msg)
	assert.False(t, isSuccess)
}

func TestModifyResponseTimestamp(t *testing.T) {
	resp := ModifyResponse{
		Code:    "",
		Message: "",
		Data:    map[string]interface{}{},
	}

	msgOK := `[{"actionName":"订单修改","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":true,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:53"},"successCount":"1"},{"actionName":"自动匹配","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"重量重算","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"表头汇总表体仓库","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"},{"actionName":"提交存量","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:09:54"},"successCount":"1"}]`
	resp.Message = msgOK

	ts := resp.Timestamp()
	assert.Equal(t, "2023-01-09 11:09:54", ts)

	msgErr := `[{"actionName":"订单修改","code":"1","externalMap":{},"failCount":"0","isExcuteAction":true,"isShowMsg":true,"sucIdAndPubts":{"1632373780662714390":"2023-01-09 11:04:34"},"successCount":"1"},{"actionName":"自动匹配","code":"1","externalMap":{},"failCount":"1","isExcuteAction":true,"isShowMsg":false,"sucIdAndPubts":{},"successCount":"0"}]`
	resp.Message = msgErr

	ts = resp.Timestamp()
	assert.Equal(t, "", ts)
}
