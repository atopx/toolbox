package middleware

import (
	"time"

	"superserver/common/logger"
	"superserver/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
server:
  name: toolbox
  env: release
  localhost: 127.0.0.1
  port: 8000
*
*/

// LoggerMiddleware 日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		logger.Info(ctx, "request", zap.String("path", ctx.Request.URL.Path))

		// 执行接口逻辑
		begin := time.Now()
		ctx.Next()
		cost := time.Since(begin)

		// response
		code := ctx.Writer.Status()
		if code < 400 || code == 404 {
			// 请求正常
			logger.Info(ctx, "response", zap.Int("code", code), zap.String("cost", cost.String()))
		} else if code < 500 {
			// 400<500 请求异常
			logger.Warn(ctx, "response", zap.String("cost", cost.String()), zap.Object(system.REPLY_ERROR_KEY,
				ctx.MustGet(system.REPLY_ERROR_KEY).(system.ReplyError)))
		} else {
			// >500 服务异常
			logger.Error(ctx, "response", zap.String("cost", cost.String()), zap.Object(system.REPLY_ERROR_KEY,
				ctx.MustGet(system.REPLY_ERROR_KEY).(system.ReplyError)))
		}
	}
}
