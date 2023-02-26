package access

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
	"superserver/proto/common_iface"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Method          string                 `json:"method"`
	Keyword         string                 `json:"keyword"`
	CreateTimeRange *common_iface.RangeI64 `json:"create_time_range"`
	UpdateTimeRange *common_iface.RangeI64 `json:"update_time_range"`
}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []AccessVo          `json:"search"`
}

type AccessVo struct {
	Id         int    `json:"id"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	Handler    string `json:"handler"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}
