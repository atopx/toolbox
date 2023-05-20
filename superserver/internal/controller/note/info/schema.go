package info

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
	"superserver/internal/model"
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

type Reply model.Note
