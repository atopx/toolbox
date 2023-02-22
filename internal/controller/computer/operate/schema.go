package operate

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
	Id      int                             `json:"id"`
	Operate *computer_iface.ComputerOperate `json:"operate"`
}

type Reply struct{}
