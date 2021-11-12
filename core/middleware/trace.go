package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从header里面拿traceid
		traceId := ctx.GetHeader("traceId")
		if traceId == "" {
			// 没有则创建
			u1 := uuid.NewV4()
			traceId = u1.String()
		}
		ctx.Set("traceId", traceId)
		ctx.Header("traceId", traceId)
		ctx.Next()
	}
}
