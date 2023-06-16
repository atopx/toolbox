package update

import (
	"cloudos/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentId"`
}

type Reply struct{}
