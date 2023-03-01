package search

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
	"superserver/internal/model/role"
	"superserver/proto/common_iface"
	"superserver/proto/user_iface"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Page   *common_iface.Pager `json:"page"`
	Filter role.Filter         `json:"filter"`
}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []RoleVo            `json:"search"`
}

type RoleVo struct {
	Id         int                   `json:"id"`
	Name       string                `json:"name"`
	Nature     user_iface.RoleNature `json:"nature"`
	Creator    string                `json:"creator"`
	Updater    string                `json:"updater"`
	CreateTime int64                 `json:"create_time"`
	UpdateTime int64                 `json:"update_time"`
}
