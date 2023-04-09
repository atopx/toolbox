package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/note/info"
	"superserver/internal/controller/note/list"
)

// NoteInfo
// @summary 笔记详情
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param body info.Params true "请求参数"
// @Response 200 object system.Response{data=info.Reply} "调用成功"
// @Router /api/note/info [get]
func (a *Api) NoteInfo(ctx *gin.Context) {
	a.Scheduler(info.NewController(ctx))
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
