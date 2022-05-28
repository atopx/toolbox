package spec

import (
	"github.com/gin-gonic/gin"
	"toolbox/internal/controller"
	"toolbox/internal/model/pixel"
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
	WidthRate  float64 `json:"width_rate" form:"width_rate"`   // 分辨率: 宽
	HeightRate float64 `json:"height_rate" form:"height_rate"` // 分辨率: 高
	Size       float64 `json:"size" form:"size"`               // 显示器尺寸
}

type Reply struct {
	PPI    float64    `json:"ppi"`    // 显示PPI
	Size   float64    `json:"size"`   // 显示器尺寸
	Height pixel.Spec `json:"height"` // 宽数据
	Width  pixel.Spec `json:"width"`  // 高数据
}
