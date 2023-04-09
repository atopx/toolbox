package api

import (
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.RouterGroup
}

func Register(engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api")}
	api.router.GET("/basic/ping", Ping)
	api.RouteUser()
	return api
}

func (a *Api) RouteUser() {
	group := a.router.Group("/user")
	group.POST("/create", a.UserCreate)
	group.POST("/login", a.UserLogin)
	group.POST("/refresh", a.UserRefresh)
	group.GET("/info", a.UserInfo)
}

func (a *Api) RouteNote() {
	group := a.router.Group("/note")
	group.GET("/info", a.NoteInfo)
	group.GET("/list", a.NoteList)
}

func (a *Api) Scheduler(ctl controller.Iface) { ctl.About(ctl.Deal()) }
