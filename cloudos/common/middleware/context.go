package middleware

import (
	"cloudos/common/consts"
	"cloudos/common/pb"
	"cloudos/common/system"
	"time"

	"cloudos/common/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ContextMiddleware 日志中间件
func (m *Middleware) ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := system.NewResponse(ctx)
		ctx.Set(consts.CK_TraceId, resp.TraceId)

		// request log
		logger.Info(ctx, "request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("client", ctx.ClientIP()),
		)

		// 执行接口逻辑
		beginTime := time.Now()
		ctx.Next()

		cost := zap.String("cost", time.Since(beginTime).String())

		if resp.Code > pb.ECode_SUCCESS {
			if resp.Code < pb.ECode_ServerInternalError {
				// error
				logger.Error(ctx, "response", cost, zap.String("error", resp.Message))
			} else {
				// warn
				logger.Warn(ctx, "response", cost, zap.String("warn", resp.Message))
			}
		} else {
			// ok
			logger.Info(ctx, "response", cost)
		}
	}
}
