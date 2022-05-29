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

// HandlerRouteRegister api注册
func HandlerRouteRegister(router *gin.Engine) {
	v1router(router.Group("/api/v1"))
}

// BaseRouteRegister 基础路由注册
func BaseRouteRegister(router *gin.Engine) {
	router.GET("/ping", base.Ping)

}

// TestRouteRegister 测试路由注册
func TestRouteRegister(router *gin.Engine) {
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
