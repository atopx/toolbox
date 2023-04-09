package api

import (
	"go.uber.org/zap"
	"superserver/common/logger"
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
		noteGroup.POST("/create", a.NoteCreate)
		noteGroup.POST("/list", a.NoteList)
	}

	// 标签
	labelGroup := a.router.Group("/label")
	{
		labelGroup.POST("/create", a.LabelCreate)
		labelGroup.POST("/list", a.LabelList)
		labelGroup.DELETE("/delete", a.LabelDelete)
	}

	// 文件夹
	folderGroup := a.router.Group("/folder")
	{
		folderGroup.POST("/create", a.FolderCreate)
		folderGroup.POST("/list", a.FolderList)
		folderGroup.DELETE("/delete", a.FolderDelete)
	}

	return a
}

type Api struct {
	router *gin.RouterGroup
}

func Register(engine *gin.Engine) *Api {
	api := &Api{router: engine.Group("/api")}
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
