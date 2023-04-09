package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/folder/create"
	"superserver/internal/controller/folder/delete"
	"superserver/internal/controller/folder/list"
)

// FolderCreate
// @summary 创建文件夹
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response{data=create.Reply} "调用成功"
// @Router /api/folder/create [post]
func (a *Api) FolderCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// FolderDelete
// @summary 删除文件夹
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/folder/delete [delete]
func (a *Api) FolderDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// FolderList
// @summary 文件夹列表
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/folder/list [post]
func (a *Api) FolderList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}
