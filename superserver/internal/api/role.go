package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/role/create"
	"superserver/internal/controller/role/delete"
	"superserver/internal/controller/role/list"
)

// RoleCreate
// @summary 创建角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response{data=create.Reply} "调用成功"
// @Router /api/role/create [post]
func (a *Api) RoleCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// RoleDelete
// @summary 删除角色
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/role/delete [delete]
func (a *Api) RoleDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// RoleList
// @summary 角色列表
// @Tags 角色
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/role/list [post]
func (a *Api) RoleList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}
