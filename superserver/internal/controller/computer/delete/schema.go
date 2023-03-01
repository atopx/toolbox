package delete

import (
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Id int `json:"id"`
}

type Reply struct{}
