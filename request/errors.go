package request

import "errors"

var (
	ErrCallYonSuiteAPIFailed = errors.New("call yonsuite api failed")
	ErrYonSuiteAPIBizError   = errors.New("call yonsuite api biz error")
	ErrAPILimit              = errors.New("reached yonsuite api limit")
)
