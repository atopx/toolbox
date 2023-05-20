package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/note/delete"
	"superserver/internal/controller/note/info"
	"superserver/internal/controller/note/list"
	"superserver/internal/controller/note/save"
)

// NoteSave
// @summary 编辑笔记
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param body save.Params true "请求参数"
// @Response 200 object system.Response{data=save.Reply} "调用成功"
// @Router /api/note/save [post]
func (a *Api) NoteSave(ctx *gin.Context) {
	a.Scheduler(save.NewController(ctx))
}

// NoteList
// @summary 笔记列表
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/note/list [post]
func (a *Api) NoteList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}

// NoteInfo
// @summary 笔记详情
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param query list.Params true "请求参数"
// @Response 200 object system.Response{data=info.Reply} "调用成功"
// @Router /api/note/info [get]
func (a *Api) NoteInfo(ctx *gin.Context) {
	a.Scheduler(info.NewController(ctx))
}

// NoteDelete
// @summary 笔记删除
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param query delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/note/delete [delete]
func (a *Api) NoteDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}
