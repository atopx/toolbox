package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const wechat_path = "wechat"

func wechat(router *gin.RouterGroup) {
	router.GET("/public", v1.WechatAuthToken)
	router.POST("/public", v1.WechatMessageListen)
	router.GET("/scheduler", v1.MiniappScheduler)
}
