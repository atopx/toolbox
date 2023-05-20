package delete

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Id *int32 `form:"id"`
}

type Reply interface{}
