package user

import (
	v1 "superserver/api/v1"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	group := router.Group("user")
	group.POST("/login", v1.UserLogin) // 登录
}
