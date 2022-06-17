package message_listen

import (
	"toolbox/internal/controller"
	"toolbox/internal/model/wechat"

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

type Param wechat.Message

// tips: xml名字不能改，会影响root tag name
type xml wechat.Message
