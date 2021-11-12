package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, data ...interface{}) {
	respJson(ctx, ErrNil, data...)
	return
}

func Error(ctx *gin.Context, errCode int, data ...interface{}) {
	respJson(ctx, errCode, data...)
	return
}

func respJson(ctx *gin.Context, errCode int, data ...interface{}) {
	var returnData interface{}
	if len(data) > 0 {
		returnData = data[0]
	} else {
		returnData = ""
	}
	ctx.AbortWithStatusJSON(http.StatusOK, Response{
		ErrCode: errCode,
		ErrMsg:  LoadErrorCode(errCode),
		Data:    returnData,
	})
	return
}

func AbortWithStatus(ctx *gin.Context, statusCode int) {
	ctx.AbortWithStatus(statusCode)
	return
}
