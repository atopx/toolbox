package controller

import (
	"net/http"
	"toolbox/common/logger"
	"toolbox/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Controller 控制器
type Controller struct {
	Context *gin.Context
	System  *system.ContextValue
	db      *gorm.DB
}

func (ctl *Controller) SetDatabase(db *gorm.DB) {
	ctl.db = db
}

func (ctl *Controller) GetDatabase() *gorm.DB {
	return ctl.db
}

// ContextLoader 上下文加载器
func (ctl *Controller) ContextLoader(ctx *gin.Context) {
	ctl.Context = ctx
	ctl.System = system.LoadContextValue(ctx)
}

// NewOkResponse 正常的response
func (ctl *Controller) NewOkResponse(code int, data interface{}) {
	ctl.Context.JSON(code, data)
}

// NewErrorResponse 异常的response
func (ctl *Controller) NewErrorResponse(code int, message string) {
	// tips: 系统级别的异常返回默认的message
	reply := system.NewReplyError(ctl.System.TraceId, code, message)
	ctl.Context.Set(system.REPLY_ERROR_KEY, reply)
	ctl.Context.JSON(code, &reply)
}

// NewRequestParam 结构化请求参数
func (ctl *Controller) NewRequestParam(param interface{}) error {
	err := ctl.Context.ShouldBind(param)
	if err != nil {
		logger.Error(ctl.Context, "bind param error", zap.Error(err))
		ctl.Context.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	return err
}
