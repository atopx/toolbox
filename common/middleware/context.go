package middleware

import (
	"toolbox/common/system"

	"github.com/gin-gonic/gin"
)

// ContextMiddleware 上下文中间件
func ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value := system.NewContextValue(ctx.ClientIP())
		value.SetUserId(0) // TODO 从token获取userId
		ctx.Set(system.CONTEXT_KEY, value)
		ctx.Next()
		system.ContextDone(value)
	}
}
