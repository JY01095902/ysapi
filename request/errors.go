package request

import "errors"

var (
	ErrCallYonSuiteAPIFailed = errors.New("call yonsuite api failed")    // 查询API未正确返回
	ErrYonSuiteAPIBizError   = errors.New("call yonsuite api biz error") // 查询API正常返回，但是结果不正确
	ErrAPILimit              = errors.New("reached yonsuite api limit")  // API限流
)
