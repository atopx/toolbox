package create

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
	ParentId int32  `json:"parent_id"`
	Source   int32  `json:"source"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}

type Reply struct{}
