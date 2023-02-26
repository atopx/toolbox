package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/permission/access"
	"superserver/internal/controller/permission/create"
	"superserver/internal/controller/permission/delete"
)

// PermissionAccess
// @summary 权限资源列表
// @Tags 角色权限
// @Accept json
// @Produce json
// @Param param body access.Params true "请求参数"
// @Response 200 access system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/permission/access [post]
func (a *Api) PermissionAccess(ctx *gin.Context) {
	a.Scheduler(access.NewController(ctx))
}

// PermissionSearch
// @summary 角色权限列表
// @Tags 角色权限
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/permission/search [post]
func (a *Api) PermissionSearch(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// PermissionCreate
// @summary 添加角色权限
// @Tags 角色权限
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/permission/create [post]
func (a *Api) PermissionCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// PermissionDelete
// @summary 删除角色权限
// @Tags 角色权限
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/permission/delete [post]
func (a *Api) PermissionDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}
