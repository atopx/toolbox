package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"superserver/common/system"
	"superserver/domain/public/ecode"
)

// OwnerMiddleware 拥有者认证
func (m *Middleware) OwnerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := system.GetRequestHeader(ctx)
		switch header.Operator {
		case system.User.Id:
			ctx.Next()
		default:
			responseHeader := system.GetResponseHeader(ctx)
			responseHeader.Code = ecode.ECode_ACCESS_Forbidden
			ctx.AbortWithStatusJSON(http.StatusForbidden, system.NewErrorResponse(responseHeader))
		}
	}
}
