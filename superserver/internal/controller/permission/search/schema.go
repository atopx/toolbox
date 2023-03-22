package search

import (
	"github.com/gin-gonic/gin"
	"superserver/common/interface/common_iface"
	"superserver/common/interface/user_iface"
	"superserver/internal/controller"
	"superserver/internal/model/permission"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Page   *common_iface.Pager `json:"page"`
	Filter permission.Filter   `json:"filter"`
}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []PermissionVo      `json:"search"`
}

type PermissionVo struct {
	Id         int       `json:"id"`
	Role       *RoleVo   `json:"role"`
	Access     *AccessVo `json:"access"`
	Creator    string    `json:"creator"`
	Updater    string    `json:"updater"`
	CreateTime int64     `json:"create_time"`
	UpdateTime int64     `json:"update_time"`
}

type RoleVo struct {
	Id     int                   `json:"id"`
	Name   string                `json:"name"`
	Nature user_iface.RoleNature `json:"nature"`
}

type AccessVo struct {
	Id      int    `json:"id"`
	Path    string `json:"path"`
	Method  string `json:"method"`
	Handler string `json:"handler"`
}
