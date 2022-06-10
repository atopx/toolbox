package miniapp

import (
	"toolbox/internal/controller"

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
	Command string `form:"command"`
}

type Reply struct {
	Status  string `json:"status"`
	Cost    string `json:"cost"`
	Message string `json:"message"`
}
