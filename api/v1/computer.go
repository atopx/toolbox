package v1

import (
	"superserver/internal/controller/computer/create"
	"superserver/internal/controller/computer/delete"
	"superserver/internal/controller/computer/operate"
	"superserver/internal/controller/computer/search"
	"superserver/internal/controller/computer/update"

	"github.com/gin-gonic/gin"
)

// ComputerCreate
// @summary 创建主机
// @Accept json
// @Produce json
// @Param param body create.Param true "请求参数"
// @Response 200 object create.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/computer/create [post]
func ComputerCreate(ctx *gin.Context) {
	if controller, err := create.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// ComputerUpdate
// @summary 更新主机
// @Accept json
// @Produce json
// @Param param body update.Param true "请求参数"
// @Response 200 object update.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/computer/update [post]
func ComputerUpdate(ctx *gin.Context) {
	if controller, err := update.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// ComputerDelete
// @summary 更新主机
// @Accept json
// @Produce json
// @Param param body delete.Param true "请求参数"
// @Response 200 object delete.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/computer/delete [delete]
func ComputerDelete(ctx *gin.Context) {
	if controller, err := delete.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// ComputerSearch
// @summary 主机查询列表
// @Accept json
// @Produce json
// @Param param body search.Param true "请求参数"
// @Response 200 object search.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/computer/search [post]
func ComputerSearch(ctx *gin.Context) {
	if controller, err := search.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// ComputerOperate
// @summary 关机
// @Accept json
// @Produce json
// @Param param body search.Param true "请求参数"
// @Response 200 object search.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/computer/operate [head]
func ComputerOperate(ctx *gin.Context) {
	if controller, err := operate.NewController(ctx); err == nil {
		controller.Deal()
	}
}
