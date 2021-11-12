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
	ErrNil   = ErrorCode(0, "成功")
	ErrParam = ErrorCode(400, "参数错误")
)

func LoadErrorCode(code int) string {
	msgVal, ok := codeMsgMap.Load(code)
	if !ok {
		msgVal = fmt.Sprintf("ErrorCode %d not define", code)
		logger.ServerLogger.Error(msgVal.(string))
	}
	return msgVal.(string)
}

func ErrorCode(code int, msg string) int {
	if _, ok := codeMsgMap.Load(code); ok {
		logger.ServerLogger.Panic(fmt.Sprintf("ErrorCode %d already exist ", code))
	}
	fmt.Println("----", code, msg)
	codeMsgMap.Store(code, msg)
	return code
}
