package api

import (
	"gorm.io/gorm"
	"superserver/common/system"
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

type Api struct {
	handler *system.Handler
	router  *gin.RouterGroup
}

func (a *Api) RouteUser() {
	group := a.router.Group("/user")
	{
		group.POST("/create", a.UserCreate)
		group.POST("/login", a.UserLogin)
		group.POST("/operate", a.ComputerOperate)
		group.DELETE("/delete", a.ComputerDelete)
	}
}

func (a *Api) RouteComputer() {
	group := a.router.Group("/computer")
	{
		group.POST("/search", a.ComputerSearch)
		group.POST("/create", a.ComputerCreate)
		group.POST("/update", a.ComputerUpdate)
		group.POST("/operate", a.ComputerOperate)
		group.DELETE("/delete", a.ComputerDelete)
	}
}

func New(db *gorm.DB, engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api")}
	api.handler = &system.Handler{Db: db}
	api.router.GET("/ping", Ping)
	api.RouteComputer()
	return api
}

func (a *Api) Scheduler(ctl controller.Iface) {
	ctl.ContextLoader()
	ctl.HandleLoader(a.handler)
	if err := ctl.RequestLoader(); err == nil {
		ctl.Deal()
	}
}
