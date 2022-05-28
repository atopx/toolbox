package message_listen

import (
	"github.com/gin-gonic/gin"
	"toolbox/internal/controller"
	"toolbox/internal/model/wechat"
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

type Param wechat.Message

type Reply wechat.Message
