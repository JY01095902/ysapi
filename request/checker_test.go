package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckResponseBody(t *testing.T) {
	body := ""
	err := checkResponseBody([]byte(body))
	assert.ErrorIs(t, err, ErrYonSuiteAPIBizError)

	body = `{
		"code": "200",
		"message": "操作成功"
	}`
	err = checkResponseBody([]byte(body))
	assert.Nil(t, err)

	body = `{
		"code": "999",
		"message": "查询失败"
	}`
	err = checkResponseBody([]byte(body))
	assert.ErrorIs(t, err, ErrYonSuiteAPIBizError)
	assert.Contains(t, err.Error(), "[999]查询失败")

	body = `{
		"code": "310046",
		"message": "API被限流"
	}`
	err = checkResponseBody([]byte(body))
	assert.ErrorIs(t, err, ErrAPILimit)
	assert.Contains(t, err.Error(), "[310046]API被限流")
}

func TestCheckResponse(t *testing.T) {
	body := `{
		"code": "200",
		"message": "操作成功"
	}`
	err := checkResponse(200, []byte(body))
	assert.Nil(t, err)

	body = `{
		"code": "310404",
		"message": "网关上没有注册此API"
	}`
	err = checkResponse(404, []byte(body))
	assert.ErrorIs(t, err, ErrCallYonSuiteAPIFailed)
	assert.Contains(t, err.Error(), "[310404]网关上没有注册此API")

	body = `{
		"code": "310046",
		"message": "API被限流"
	}`
	err = checkResponse(429, []byte(body))
	assert.ErrorIs(t, err, ErrAPILimit)
	assert.Contains(t, err.Error(), "[310046]API被限流")

	body = `调用API失败`
	err = checkResponse(500, []byte(body))
	assert.ErrorIs(t, err, ErrCallYonSuiteAPIFailed)
	assert.Contains(t, err.Error(), "调用API失败")
}
