package controller

import (
	"cloudos/common/pb"
	"cloudos/common/system"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Controller 控制器
type Controller struct {
	context *gin.Context
	resp    *system.Response
	Params  any
	Handler *system.Handler
	error   error
}

func New(ctx *gin.Context, params any) *Controller {
	err := ctx.ShouldBind(params)
	return &Controller{
		context: ctx,
		Params:  params,
		Handler: system.GetHandler(),
		resp:    system.GetResponse(ctx),
		error:   err,
	}
}

func (ctl *Controller) Deal() (any, pb.ECode) {
	return nil, pb.ECode_Unimplemented
}

func (ctl *Controller) About(data any, code pb.ECode) {
	ctl.resp.Code = code
	ctl.resp.Message = code.String()
	ctl.resp.Data = data
	ctl.context.JSON(http.StatusOK, ctl.resp)
}

func (ctl *Controller) Error() error { return ctl.error }

func (ctl *Controller) Context() *gin.Context { return ctl.context }

type Iface interface {
	Deal() (any, pb.ECode)
	Error() error
	Context() *gin.Context
	About(any, pb.ECode)
}
