package trans

import (
	"toolbox/internal/controller"

	"github.com/gin-gonic/gin"
)

type Param struct {
	JsonStr string `form:"json_str"` // json字符串
}

type Reply map[string]interface{}

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
