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
	api.router.GET("/basic/ping", Ping)
	api.RouteComputer()
	api.RouteUser()
	api.RouteRole()
	api.RouteNovel()
	api.RoutePermission()
	return api
}

func (a *Api) RouteUser() {
	group := a.router.Group("/user")
	{
		group.POST("/create", a.UserCreate)
		group.POST("/login", a.UserLogin)
		group.POST("/refresh", a.UserRefresh)
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

func (a *Api) RouteNovel() {
	group := a.router.Group("/novel")
	{
		group.POST("/create", a.NovelCreate)
	}
}

func (a *Api) RoutePermission() {
	group := a.router.Group("/permission")
	{
		group.POST("/access", a.PermissionAccess)
		group.POST("/create", a.PermissionCreate)
		group.POST("/search", a.PermissionSearch)
		group.DELETE("/delete", a.PermissionDelete)
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
		group.POST("/scheduler", a.ComputerScheduler)
	}
}

func (a *Api) Scheduler(ctl controller.Iface) {
	ctl.ContextLoader()
	ctl.HandleLoader(a.handler)
	if err := ctl.RequestParamsLoader(); err == nil {
		ctl.Deal()
	}
}
