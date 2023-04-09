package controller

import (
	"net/http"
	"superserver/common/consts"
	"superserver/common/system"
	"superserver/domain/public/common"
	"superserver/domain/public/ecode"

	"github.com/gin-gonic/gin"
)

// Controller 控制器
type Controller struct {
	context *gin.Context
	header  *common.Header
	Params  any
	error   error
}

func New(ctx *gin.Context, params any) *Controller {
	err := ctx.ShouldBind(params)
	return &Controller{
		context: ctx,
		Params:  params,
		header:  system.GetRequestHeader(ctx),
		error:   err,
	}
}

func (ctl *Controller) NewServiceHeader() *common.Header {
	return &common.Header{
		TraceId:  ctl.header.TraceId,
		Source:   consts.ServiceName,
		Operator: ctl.header.Operator,
	}
}

func (ctl *Controller) Deal() (any, ecode.ECode) {
	return nil, ecode.ECode_SYSTEM_ERROR_Unimplemented
}

func (ctl *Controller) About(data any, code ecode.ECode) {
	header := system.GetResponseHeader(ctl.context)
	header.Code = code
	ctl.context.JSON(http.StatusOK, system.NewResponse(header, data))
}

func (ctl *Controller) Error() error { return ctl.error }

func (ctl *Controller) Context() *gin.Context { return ctl.context }

func (ctl *Controller) Header() *common.Header { return ctl.header }

type Iface interface {
	Deal() (any, ecode.ECode)
	Error() error
	Header() *common.Header
	Context() *gin.Context
	About(any, ecode.ECode)
}
