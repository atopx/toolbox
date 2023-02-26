package access

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
	"superserver/internal/model/access"
	"superserver/proto/common_iface"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Page   *common_iface.Pager `json:"page"`
	Filter access.Filter       `json:"filter"`
}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []AccessVo          `json:"search"`
}

type AccessVo struct {
	Id         int                       `json:"id"`
	Path       string                    `json:"path"`
	Method     common_iface.AccessMethod `json:"method"`
	Handler    string                    `json:"handler"`
	CreateTime int64                     `json:"create_time"`
	UpdateTime int64                     `json:"update_time"`
}
