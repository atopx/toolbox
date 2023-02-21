package middleware

import (
	"superserver/common/system"

	"github.com/gin-gonic/gin"
)

// ContextMiddleware 上下文中间件
func ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// set request context
		trace := system.NewTrace()
		ctx.Set(system.TRACE_KEY, trace)
		ctx.Next()
	}
}
