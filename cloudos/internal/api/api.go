package api

import (
	"cloudos/common/logger"
	"cloudos/common/middleware"
	"cloudos/common/pb"
	"cloudos/internal/controller"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func (a *Api) Route() *Api {
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
	api.router.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })

	userGroup := api.router.Group("/user")
	{
		userGroup.POST("/login", api.UserLogin)
	}

	folderGroup := api.router.Group("/folder")
	{
		folderGroup.POST("/list", api.FolderList)
		folderGroup.POST("/create", api.FolderCreate)
		folderGroup.POST("/update", api.FolderUpdate)
		folderGroup.POST("/delete", api.FolderDelete)
	}

	return api.Route()
}

func (a *Api) Scheduler(ctl controller.Iface) {
	if err := ctl.Error(); err != nil {
		logger.Warn(ctl.Context(), "bind param error", zap.Error(err))
		ctl.About(nil, pb.ECode_BadRequest)
		return
	}
	ctl.About(ctl.Deal())
}
