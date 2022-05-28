package auth_token

import (
	"github.com/gin-gonic/gin"
	"toolbox/internal/controller"
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
	Signature string `json:"Signature" form:"signature"` // 签名
	Timestamp string `json:"Timestamp" form:"timestamp"` // 时间戳
	Nonce     string `json:"Nonce" form:"nonce"`         // 随机字符串
	Echostr   string `json:"Echostr" form:"echostr"`     // 回填字符串
}
