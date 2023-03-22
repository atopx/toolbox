package controller

import (
	"net/http"
	"superserver/common/interface/ecode_iface"

	"superserver/common/logger"
	"superserver/common/system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Controller 控制器
type Controller struct {
	Context *gin.Context
	db      *gorm.DB
	chain   *system.ChainMessage
	Params  any
}

func (ctl *Controller) GetDatabase() *gorm.DB {
	return ctl.db
}

func (ctl *Controller) GetAuthTime() int64 {
	return ctl.chain.AuthTime
}

// ContextLoader 上下文加载器
func (ctl *Controller) ContextLoader() {
	ctl.chain = system.GetChainMessage(ctl.Context)
}

func (ctl *Controller) GetOperator() int {
	return ctl.chain.UserId
}

// HandleLoader 句柄加载器
func (ctl *Controller) HandleLoader(handler *system.Handler) {
	ctl.db = handler.Db
}

// RequestParamsLoader 参数加载器
func (ctl *Controller) RequestParamsLoader() error {
	err := ctl.Context.ShouldBind(ctl.Params)
	if err != nil {
		logger.Error(ctl.Context, "bind param error", zap.Error(err))
		ctl.Context.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	return err
}

// NewOkResponse 正常的response
func (ctl *Controller) NewOkResponse(data any) {
	ctl.chain.WriteNormal(data)
	ctl.Context.JSON(http.StatusOK, ctl.chain)
}

// NewErrorResponse 异常的response
func (ctl *Controller) NewErrorResponse(ecode ecode_iface.ECode, message string) {
	ctl.chain.WriteAbnormal(ecode_iface.ECode_BAD_REQUEST, message)
	var code int
	if ecode >= ecode_iface.ECode_SYSTEM_ERROR {
		code = http.StatusInternalServerError
	} else {
		code = http.StatusBadRequest
	}
	ctl.Context.JSON(code, ctl.chain)
}

type Iface interface {
	ContextLoader()
	HandleLoader(*system.Handler)
	RequestParamsLoader() error
	Deal()
}
