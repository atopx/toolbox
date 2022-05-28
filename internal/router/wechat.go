package router

import (
	"github.com/gin-gonic/gin"
	v1 "toolbox/api/v1"
)

const wechat_path = "wechat"

func wechat(router *gin.RouterGroup) {
	router.GET("/public", v1.WechatAuthToken)
	router.POST("/public", v1.WechatMessageListen)
}
