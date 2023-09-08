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

type Api struct {
	router    *gin.RouterGroup
	middle    *middleware.Middleware
	AccessMap map[string]pb.UserRole
}

func Register(engine *gin.Engine) *Api {
	api := &Api{router: &engine.RouterGroup, middle: middleware.New()}
	api.router.Use(
		api.middle.CorsMiddleware(),
		api.middle.RecoverMiddleware(),
		api.middle.ContextMiddleware(),
		api.middle.AuthMiddleware(),
	)
	api.router.GET("/ping", func(ctx *gin.Context) { ctx.String(http.StatusOK, "pong") })

	// user
	api.route(http.MethodPost, "/api/v1/user/login", api.UserLogin, pb.UserRole_GUEST)
	api.route(http.MethodPost, "/api/v1/user/refresh", api.UserRefresh, pb.UserRole_GUEST)

	// note
	api.route(http.MethodPost, "/api/v1/note/list", api.NoteList, pb.UserRole_USER)
	api.route(http.MethodPost, "/api/v1/note/save", api.NoteSave, pb.UserRole_USER)
	api.route(http.MethodPost, "/api/v1/note/delete", api.NoteDelete, pb.UserRole_USER)
	api.route(http.MethodPost, "/api/v1/note/addtag", api.NoteAddtag, pb.UserRole_USER)
	api.route(http.MethodPost, "/api/v1/note/deltag", api.NoteDeltag, pb.UserRole_USER)
	api.route(http.MethodGet, "/api/v1/note/topics", api.NoteTopics, pb.UserRole_USER)
	api.route(http.MethodGet, "/api/v1/note/info", api.NoteInfo, pb.UserRole_USER)

	return api
}

func (a *Api) route(method, path string, handler gin.HandlerFunc, role pb.UserRole) {
	a.middle.SetAccess(method, path, role)
	a.router.Handle(method, path, handler)
}

func (a *Api) Scheduler(ctl controller.Iface) {
	if err := ctl.Error(); err != nil {
		logger.Warn(ctl.Context(), "bind param error", zap.Error(err))
		ctl.About(nil, pb.ECode_BadRequest)
		return
	}
	ctl.About(ctl.Deal())
}
