package tradeorder

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseResponseIsSuccessed(t *testing.T) {
	msgAllOK := `{
		"code": "200",
		"message": "操作成功",
		"data": [
			{
				"code": "1",
				"isShowMsg": true,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321224000
				},
				"successCount": "1",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321224000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消顺丰WMS出库单"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321224000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "提交存量"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321224000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消奇门单据"
			},
			{
				"code": "1",
				"isShowMsg": true,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183727": 1675321225000
				},
				"successCount": "1",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183727": 1675321225000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消顺丰WMS出库单"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183727": 1675321225000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "提交存量"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183727": 1675321225000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消奇门单据"
			}
		]
	}`
	resp := CloseResponse{}
	_ = json.Unmarshal([]byte(msgAllOK), &resp)

	for _, id := range []string{"1648795880186183721", "1648795880186183727"} {
		isSuccess, msg := resp.IsSuccessed(id)
		assert.Empty(t, msg)
		assert.True(t, isSuccess)
	}

	msgAllErr := `{
		"code": "200",
		"message": "操作成功",
		"data": [
			{
				"exceptionMsg": "没有需要操作的单据",
				"code": "1",
				"isShowMsg": true,
				"failCount": "1",
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			},
			{
				"exceptionMsg": "没有需要操作的单据",
				"code": "1",
				"isShowMsg": true,
				"failCount": "1",
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			}
		]
	}`
	_ = json.Unmarshal([]byte(msgAllErr), &resp)
	for _, id := range []string{"1648795880186183721", "1648795880186183727"} {
		isSuccess, msg := resp.IsSuccessed(id)
		assert.NotEmpty(t, msg)
		assert.False(t, isSuccess)
	}

	msgOK1Err1 := `{
		"code": "200",
		"message": "操作成功",
		"data": [
			{
				"code": "1",
				"isShowMsg": true,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321913000
				},
				"successCount": "1",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321913000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消顺丰WMS出库单"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321913000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "提交存量"
			},
			{
				"code": "1",
				"isShowMsg": false,
				"externalMap": {},
				"failCount": "0",
				"sucIdAndPubts": {
					"1648795880186183721": 1675321913000
				},
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "取消奇门单据"
			},
			{
				"exceptionMsg": "没有需要操作的单据",
				"code": "1",
				"isShowMsg": true,
				"failCount": "1",
				"successCount": "0",
				"isExcuteAction": true,
				"actionName": "订单关闭"
			}
		]
	}`

	_ = json.Unmarshal([]byte(msgOK1Err1), &resp)

	isSuccess, msg := resp.IsSuccessed("1648795880186183721")
	assert.Empty(t, msg)
	assert.True(t, isSuccess)

	isSuccess, msg = resp.IsSuccessed("1648795880186183727")
	assert.NotEmpty(t, msg)
	assert.False(t, isSuccess)
}
