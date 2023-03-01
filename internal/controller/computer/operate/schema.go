package operate

import (
	"superserver/internal/controller"
	"superserver/proto/computer_iface"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Id      int                             `json:"id"`
	Operate *computer_iface.ComputerOperate `json:"operate"`
}

type Reply struct{}
