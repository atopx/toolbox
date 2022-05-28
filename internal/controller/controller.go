package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"toolbox/common/logger"
	"toolbox/common/system"
)

// Controller 控制器
type Controller struct {
	Context *gin.Context
	System  *system.ContextValue
	db      *gorm.DB
}

// ContextLoader 上下文加载器
func (ctl *Controller) ContextLoader(ctx *gin.Context) {
	ctl.Context = ctx
	ctl.System = system.LoadContextValue(ctx)
}

// NewOkResponse 正常的response
func (ctl *Controller) NewOkResponse(code int, data interface{}) {
	logger.Info(ctl.Context, http.StatusText(code))
	ctl.Context.JSON(code, data)
}

// NewWarnResponse 请求异常的response
func (ctl *Controller) NewWarnResponse(code int, message string) {
	general := http.StatusText(code)
	if message == "" {
		message = general
	}
	err := system.ReplyError{Code: code, Message: message}
	logger.Warn(ctl.Context, general, zap.Object("response", err))
	ctl.Context.JSON(code, &err)
}

// NewErrorResponse 服务异常的response
func (ctl *Controller) NewErrorResponse(code int, message string) {
	// tips: 系统级别的异常返回默认的message
	general := http.StatusText(code)
	if message == "" {
		message = general
	}
	err := system.ReplyError{Code: code, Message: general}
	logger.Error(ctl.Context, message, zap.Object("response", err))
	ctl.Context.JSON(code, &err)
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
