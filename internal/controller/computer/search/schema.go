package search

import (
	"superserver/internal/controller"
	"superserver/internal/model/computer"
	"superserver/proto/common_iface"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Page   *common_iface.Pager `json:"page"`
	Filter computer.Filter     `json:"filter"`
}

type Reply struct {
	Page *common_iface.Pager `json:"page"`
	List []ComputerVo        `json:"list"`
}

type ComputerVo struct {
	Id          int    `json:"id"`           // ID
	Name        string `json:"name"`         // 名称
	Username    string `json:"username"`     // 用户名
	Password    string `json:"password"`     // 用户密码
	LanHostname string `json:"lan_hostname"` // 局域网地址
	WanHostname string `json:"wan_hostname"` // 广域网地址
	Address     string `json:"address"`      // 物理地址
	PowerStatus string `json:"power_status"` // 电源状态
	Creator     string `json:"creator"`      // 创建人
	Updator     string `json:"updator"`      // 更新人
	CreateTime  int64  `json:"create_time"`  // 创建时间
	UpdateTime  int64  `json:"update_time"`  // 更新时间
	ScanTime    int64  `json:"scan_time"`    // 最后一次扫描时间
}
