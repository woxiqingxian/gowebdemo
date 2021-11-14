package response

import (
	"fmt"
	"gowebdemo/core/logger"
	"sync"
)

var (
	codeMsgMap sync.Map
)

// 通用错误码
var (
	ErrNil = ErrorCode(0, "成功")
)

func LoadErrorCode(code int) string {
	msgVal, ok := codeMsgMap.Load(code)
	if !ok {
		msgVal = fmt.Sprintf("ErrorCode %d not define", code)
		logger.ServerLog().Error(msgVal.(string))
	}
	return msgVal.(string)
}

func ErrorCode(code int, msg string) int {
	if _, ok := codeMsgMap.Load(code); ok {
		logger.ServerLog().Panic(fmt.Sprintf("ErrorCode %d already exist ", code))
	}
	codeMsgMap.Store(code, msg)
	return code
}
