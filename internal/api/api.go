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

func New(db *gorm.DB, engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api")}
	api.handler = &system.Handler{Db: db}
	api.router.GET("/ping", Ping)
	api.RouteComputer()
	api.RouteUser()
	api.RouteRole()
	return api
}

func (a *Api) RouteUser() {
	group := a.router.Group("/user")
	{
		group.POST("/create", a.UserCreate)
		group.POST("/login", a.UserLogin)
	}
}

func (a *Api) RouteRole() {
	group := a.router.Group("/role")
	{
		group.POST("/create", a.RoleCreate)
		group.POST("/update", a.RoleUpdate)
		group.POST("/search", a.RoleSearch)
		group.DELETE("/delete", a.RoleDelete)
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

func (a *Api) Scheduler(ctl controller.Iface) {
	ctl.ContextLoader()
	ctl.HandleLoader(a.handler)
	if err := ctl.RequestLoader(); err == nil {
		ctl.Deal()
	}
}
