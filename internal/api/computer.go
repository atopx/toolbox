package api

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
// Tags 主机
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/computer/create [post]
func (a *Api) ComputerCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// ComputerUpdate
// @summary 更新主机
// Tags 主机
// @Accept json
// @Produce json
// @Param param body update.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/computer/update [post]
func (a *Api) ComputerUpdate(ctx *gin.Context) {
	a.Scheduler(update.NewController(ctx))
}

// ComputerDelete
// @summary 删除主机
// Tags 主机
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.ChainMessage{data=delete.Reply} "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/computer/delete [delete]
func (a *Api) ComputerDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// ComputerSearch
// @summary 查询主机
// Tags 主机
// @Accept json
// @Produce json
// @Param param body search.Params true "请求参数"
// @Response 200 object system.ChainMessage{data=search.Reply} "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/computer/search [post]
func (a *Api) ComputerSearch(ctx *gin.Context) {
	a.Scheduler(search.NewController(ctx))
}

// ComputerOperate
// @summary 操作主机
// Tags 主机
// @Accept json
// @Produce json
// @Param param body operate.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/computer/operate [post]
func (a *Api) ComputerOperate(ctx *gin.Context) {
	a.Scheduler(operate.NewController(ctx))
}
