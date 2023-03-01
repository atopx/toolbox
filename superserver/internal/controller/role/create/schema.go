package create

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Name string `json:"name"`
}

type Reply struct{}
