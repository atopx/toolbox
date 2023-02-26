package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/role/create"
	"superserver/internal/controller/role/delete"
	"superserver/internal/controller/role/search"
	"superserver/internal/controller/role/udpate"
)

// RoleCreate
// @summary 创建角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/role/create [post]
func (a *Api) RoleCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// RoleUpdate
// @summary 更新角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body update.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/role/update [post]
func (a *Api) RoleUpdate(ctx *gin.Context) {
	a.Scheduler(update.NewController(ctx))
}

// RoleDelete
// @summary 删除角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/role/delete [delete]
func (a *Api) RoleDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// RoleSearch
// @summary 角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body search.Params true "请求参数"
// @Response 200 object system.ChainMessage{data=search.Reply} "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/role/search [post]
func (a *Api) RoleSearch(ctx *gin.Context) {
	a.Scheduler(search.NewController(ctx))
}
