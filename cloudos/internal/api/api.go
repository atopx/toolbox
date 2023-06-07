package api

import (
	"cloudos/common/system"
	"cloudos/internal/controller"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type Api struct {
	handler *system.Handler
	router  *gin.RouterGroup
}

func New(db *gorm.DB, engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api")}
	api.handler = &system.Handler{Db: db}
	api.router.GET("/basic/ping", Ping)
	api.RouteUser()
	return api
}

func (a *Api) RouteUser() {
	group := a.router.Group("/user")
	{
		group.POST("/create", a.UserCreate)
		group.POST("/login", a.UserLogin)
		group.POST("/refresh", a.UserRefresh)
		group.GET("/info", a.UserInfo)
	}
}

func (a *Api) Scheduler(ctl controller.Iface) {
	ctl.ContextLoader()
	ctl.HandleLoader(a.handler)
	if err := ctl.RequestParamsLoader(); err == nil {
		ctl.Deal()
	}
}
