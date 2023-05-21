package api

import (
	"go.uber.org/zap"
	"superserver/common/logger"
	"superserver/common/middleware"
	"superserver/domain/public/ecode"
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

func (a *Api) Route() *Api {

	// 用户
	userGroup := a.router.Group("/user")
	{
		userGroup.POST("/create", a.UserCreate)
		userGroup.POST("/login", a.UserLogin)
		userGroup.POST("/refresh", a.UserRefresh)
		userGroup.GET("/info", a.UserInfo)
	}

	// 笔记
	noteGroup := a.router.Group("/note")
	{
		noteGroup.POST("/save", a.NoteSave)
		noteGroup.POST("/list", a.NoteList)
		noteGroup.GET("/info", a.NoteInfo)
		noteGroup.DELETE("/delete", a.NoteDelete)
		noteGroup.DELETE("/deltag", a.NoteDeltag)
		noteGroup.POST("/addtag", a.NoteAddtag)
	}

	// 标签
	labelGroup := a.router.Group("/label")
	{
		labelGroup.POST("/create", a.LabelCreate)
		labelGroup.POST("/list", a.LabelList)
		labelGroup.DELETE("/delete", a.LabelDelete)
	}

	// 角色
	roleGroup := a.router.Group("/role", a.middle.OwnerMiddleware())
	{
		roleGroup.POST("/create", a.RoleCreate)
		roleGroup.POST("/list", a.RoleList)
		roleGroup.DELETE("/delete", a.RoleDelete)
	}

	// 资源
	//accessGroup := a.router.Group("/access", a.middle.OwnerMiddleware())
	//{
	//	accessGroup.POST("/create", a.AccessCreate)
	//	accessGroup.POST("/list", a.AccessList)
	//	accessGroup.DELETE("/delete", a.AccessDelete)
	//}

	// 权限
	//permissionGroup := a.router.Group("/permission", a.middle.OwnerMiddleware())
	//{
	//	permissionGroup.POST("/create", a.PermissionCreate)
	//	permissionGroup.POST("/list", a.PermissionList)
	//	permissionGroup.DELETE("/delete", a.PermissionDelete)
	//}

	return a
}

type Api struct {
	router *gin.RouterGroup
	middle *middleware.Middleware
}

func Register(engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api"), middle: middleware.New()}
	api.router.Use(
		api.middle.CorsMiddleware(),
		api.middle.RecoverMiddleware(),
		api.middle.ContextMiddleware(),
		api.middle.AuthMiddleware(),
	)
	api.router.GET("/basic/ping", api.Ping)
	return api.Route()
}

func (a *Api) Scheduler(ctl controller.Iface) {
	if err := ctl.Error(); err != nil {
		logger.Warn(ctl.Context(), "bind param error", zap.Error(err))
		ctl.About(nil, ecode.ECode_BAD_REQUEST)
	} else {
		ctl.About(ctl.Deal())
	}
}
