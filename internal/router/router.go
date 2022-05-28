package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"toolbox/api/base"
	"toolbox/docs"
)

func v1router(router *gin.RouterGroup) {
	wechat(router.Group(wechat_path))
	pixel(router.Group(pixel_path))
}

// HandlerRegister api注册
func HandlerRegister(router *gin.Engine) {
	v1router(router.Group("/v1"))
}

// BaseRegister 基础路由注册
func BaseRegister(router *gin.Engine) {
	router.GET("/ping", base.Ping)
	switch gin.Mode() {
	case gin.DebugMode:
		docs.SwaggerInfo.Host = fmt.Sprintf(
			"%s:%d",
			viper.GetString("server.localhost"),
			viper.GetInt("server.port"),
		)
		router.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
	}
}
