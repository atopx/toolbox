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
	chain   *system.ChainContext
}

func (ctl *Controller) SetDatabase(db *gorm.DB) {
	ctl.db = db.WithContext(ctl.Context)
}

func (ctl *Controller) GetDatabase() *gorm.DB {
	return ctl.db
}

// ContextLoader 上下文加载器
func (ctl *Controller) ContextLoader(ctx *gin.Context) {
	ctl.Context = ctx
	message, _ := ctx.Get(system.ChainContextKey)
	ctl.chain = message.(*system.ChainContext)
}

// NewOkResponse 正常的response
func (ctl *Controller) NewOkResponse(code int, data interface{}) {
	ctl.chain.WriteNormal(data)
	ctl.Context.JSON(code, ctl.chain)
}

// NewErrorResponse 异常的response
func (ctl *Controller) NewErrorResponse(code int, message string) {
	ctl.chain.WriteAbnomal(system.CHAIN_BAD, message)
	ctl.Context.JSON(code, ctl.chain)
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
