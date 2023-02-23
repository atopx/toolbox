package controller

import (
	"net/http"

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

// ContextLoader 上下文加载器
func (ctl *Controller) ContextLoader() {
	ctl.chain = system.GetChainMessage(ctl.Context)
}

// HandleLoader 句柄加载器
func (ctl *Controller) HandleLoader(handler *system.Handler) {
	ctl.db = handler.Db
}

// ParamLoader 参数加载器
func (ctl *Controller) RequestLoader() error {
	err := ctl.Context.ShouldBind(ctl.Params)
	if err != nil {
		logger.Error(ctl.Context, "bind param error", zap.Error(err))
		ctl.Context.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
	return err
}

// NewOkResponse 正常的response
func (ctl *Controller) NewOkResponse(code int, data any) {
	ctl.chain.WriteNormal(data)
	ctl.Context.JSON(code, ctl.chain)
}

// NewErrorResponse 异常的response
func (ctl *Controller) NewErrorResponse(code int, message string) {
	ctl.chain.WriteAbnomal(system.CHAIN_BAD, message)
	ctl.Context.JSON(code, ctl.chain)
}

type Iface interface {
	ContextLoader()
	HandleLoader(*system.Handler)
	RequestLoader() error
	Deal()
}
