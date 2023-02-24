package middleware

import (
	"net/http"
	"strings"
	"superserver/common/system"
	"superserver/internal/model/access"
	"superserver/internal/model/permission"
	"superserver/proto/user_iface"

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
		claims, err := system.GetClaims(token)
		if err != nil {
			chain.Message = err.Error()
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

		// 系统管理员直接放行 & TODO common_iface.AccessStatus_ANY...
		if user_iface.UserRole(claims.RoleId) != user_iface.UserRole_USER_ROLE_SYSTEM && acces.Status != 1 {
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
