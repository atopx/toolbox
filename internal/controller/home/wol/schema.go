package wol

import (
	"superserver/internal/controller"

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
	MacAddr string `form:"mac_addr"` // mac地址
}

type Reply struct {
	Message string `json:"message"`
}
