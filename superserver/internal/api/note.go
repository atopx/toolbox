package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/note/create"
	"superserver/internal/controller/note/list"
)

// NoteCreate
// @summary 新建笔记
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response{data=create.Reply} "调用成功"
// @Router /api/note/create [post]
func (a *Api) NoteCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
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
