package controller

import (
	"net/http"
	"superserver/common/consts"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Controller 控制器
type Controller struct {
	Context *gin.Context
	Header  *common.Header
	Params  any
}

func New(ctx *gin.Context, params any) *Controller {
	err := ctx.ShouldBind(params)
	if err != nil {
		logger.Error(ctx, "bind param error", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		ctx.Abort()
	}
	return &Controller{
		Context: ctx,
		Params:  params,
		Header:  system.GetRequestHeader(ctx),
	}
}

func (ctl *Controller) NewServiceHeader() *common.Header {
	return &common.Header{
		TraceId:  ctl.Header.TraceId,
		Source:   consts.ServiceName,
		Operator: ctl.Header.Operator,
	}
}

func (ctl *Controller) Deal() (any, ecode.ECode) {
	return nil, ecode.ECode_SYSTEM_ERROR_Unimplemented
}

func (ctl *Controller) About(data any, code ecode.ECode) {
	header := system.GetResponseHeader(ctl.Context)
	header.Code = code
	ctl.Context.JSON(http.StatusOK, system.NewResponse(header, data))
}

type Iface interface {
	Deal() (any, ecode.ECode)
	About(any, ecode.ECode)
}
