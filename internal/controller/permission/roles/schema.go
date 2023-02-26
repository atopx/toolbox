package roles

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

type Params struct{}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []RoleVo            `json:"list"`
}

type RoleVo struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
	Name  string `json:"name"`
}
