package api

import (
	"cloudos/internal/controller/folder/create"
	"cloudos/internal/controller/folder/delete"
	"cloudos/internal/controller/folder/list"
	"cloudos/internal/controller/folder/update"

	"github.com/gin-gonic/gin"
)

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

// FolderCreate
// @summary 新建文件夹
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response{data=create.Reply} "调用成功"
// @Router /api/folder/create [post]
func (a *Api) FolderCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// FolderUpdate
// @summary 更新文件夹
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body update.Params true "请求参数"
// @Response 200 object system.Response{data=update.Reply} "调用成功"
// @Router /api/folder/update [post]
func (a *Api) FolderUpdate(ctx *gin.Context) {
	a.Scheduler(update.NewController(ctx))
}

// FolderDelete
// @summary 删除文件夹
// @Tags 文件夹
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/folder/delete [post]
func (a *Api) FolderDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}
