package uconfig

import "gowebdemo/core/response"

// 通用错误码 放在了
// core/response/ecode.go
var (
	ErrNil = response.ErrNil
)

// 业务错误码
var (
	ErrDemo = response.ErrorCode(10001, "业务demo错误")
)
