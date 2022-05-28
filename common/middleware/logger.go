package middleware

import "github.com/gin-gonic/gin"

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO 待实现
		ctx.Next()
	}
}
