package middleware

import (
	"superserver/common/system"
	"superserver/common/utils"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"
	"time"

	"superserver/common/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ContextMiddleware 日志中间件
func (m *Middleware) ContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// new context
		source := ctx.GetHeader("source")
		system.SetRequestHeader(ctx, &common.Header{
			TraceId:  utils.NewTraceId(),
			Source:   source,
			Operator: 0,
		})
		responseHeader := &common.ReplyHeader{
			TraceId: utils.NewTraceId(),
			Code:    ecode.ECode_SUCCESS,
			Message: "",
		}
		system.SetResponseHeader(ctx, responseHeader)

		// request log
		logger.Info(ctx, "request",
			zap.String("method", ctx.Request.Method),
			zap.String("path", ctx.Request.URL.Path),
			zap.String("source", source),
			zap.String("client", ctx.ClientIP()),
		)

		// 执行接口逻辑
		beginTime := time.Now()
		ctx.Next()

		cost := zap.String("cost", time.Since(beginTime).String())
		if responseHeader.Code < ecode.ECode_BAD_REQUEST {
			logger.Info(ctx, "response", cost)
		} else {
			if responseHeader.Code < ecode.ECode_SYSTEM_INTERNAL_ERROR {
				logger.Warn(ctx, "response", cost, zap.String("warn", responseHeader.Message))
			} else {
				logger.Error(ctx, "response", cost, zap.String("error", responseHeader.Message))
			}
		}
	}
}
