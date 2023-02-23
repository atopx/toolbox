package create

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
	Name         string                                `json:"name"`
	Username     string                                `json:"username"`
	Password     string                                `json:"password"`
	LanHostname  string                                `json:"lan_hostname"`
	WanHostname  string                                `json:"wan_hostname"`
	Address      string                                `json:"address"`
	DefaultPorts []computer_iface.ComputerPortProtocol `json:"default_ports"`
}

type Reply struct{}
