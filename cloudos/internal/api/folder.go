package api

import (
	"cloudos/internal/controller/folder/list"

	"github.com/gin-gonic/gin"
)

// FolderList
// @summary 文件夹列表
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/folder/list [post]
func (a *Api) FolderList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}
