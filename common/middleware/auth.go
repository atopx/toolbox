package middleware

import (
	"go.uber.org/zap"
	"net/http"
	"strings"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/internal/model/access"
	"superserver/internal/model/permission"
	"superserver/proto/common_iface"
	"superserver/proto/ecode_iface"
	"superserver/proto/user_iface"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		chain := system.GetChainMessage(ctx)
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			chain.Message = "账户未登录"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
			return
		}
		claims, err := system.UnSignClaims(token)
		if err != nil {
			logger.Error(ctx, "auth system.UnSignClaims failed", zap.Error(err))
			chain.Message = ecode_iface.ECode_AUTH_INVALID.String()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
			return
		}
		if claims.ExpiresAt.Before(time.Now().Local()) {
			chain.Message = ecode_iface.ECode_AUTH_EXPIRED.String()
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
			return
		}

		// 查询access是否是游客权限
		value, ok := access.AccessMap.Load(ctx.Request.RequestURI)
		if !ok {
			chain.Message = "404"
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
			return
		}
		acces := value.(*access.Access)

		// 系统管理员直接放行或匿名接口
		if user_iface.UserRole(claims.RoleId) != user_iface.UserRole_USER_ROLE_SYSTEM && acces.Status != common_iface.AccessStatus_ACCESS_ANONYMOUS {
			if !permission.NewDao(m.db).Inspector(claims.RoleId, acces.Id) {
				chain.Message = "无权限"
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, chain)
				return
			}
		}
		chain.UserId = claims.UserId
		ctx.Next()
	}
}
