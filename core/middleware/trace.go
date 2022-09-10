package middleware

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// Trace 中间件 for 链路追踪
func Trace() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从header里面拿traceid
		traceID := ctx.GetHeader("traceID")
		if traceID == "" {
			// 没有则创建
			u1 := uuid.NewV4()
			traceID = u1.String()
		}
		ctx.Set("traceId", traceID)
		ctx.Header("traceId", traceID)
		ctx.Next()
	}
}
