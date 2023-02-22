package create

import (
	"superserver/internal/controller"
	"superserver/proto/computer_iface"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	param *Param
	controller.Controller
}

func NewController(ctx *gin.Context) (*Controller, error) {
	ctl := Controller{param: new(Param)}
	ctl.ContextLoader(ctx)
	if err := ctl.NewRequestParam(ctl.param); err != nil {
		return nil, err
	}
	return &ctl, nil
}

type Param struct {
	Name         string                                `json:"name"`
	Username     string                                `json:"username"`
	Password     string                                `json:"password"`
	LanHostname  string                                `json:"lan_hostname"`
	WanHostname  string                                `json:"wan_hostname"`
	Address      string                                `json:"address"`
	DefaultPorts []computer_iface.ComputerPortProtocol `json:"default_ports"`
}

type Reply struct{}
