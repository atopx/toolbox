package server

import (
	"fmt"
	"superserver/api/base"
	"superserver/docs"
	"superserver/internal/controller/user"

	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func (srv *Server) routeV1Group() {
	v1group := srv.engine.Group("/api/v1")

	// user router
	user.Route(v1group)

}

func (srv *Server) Route() {
	srv.engine.GET("/ping", base.Ping)
	srv.routeV1Group()
	if gin.Mode() == gin.DebugMode {
		srv.routeSwagger()
	}
}

// TestRouteRegister 测试路由注册
func (srv *Server) routeSwagger() {
	docs.SwaggerInfo.Host = fmt.Sprintf(
		"%s:%d",
		viper.GetString("server.localhost"),
		viper.GetInt("server.port"),
	)
	srv.engine.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
}
