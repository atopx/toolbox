package api

import (
	"cloudos/internal/controller/note/addtag"
	"cloudos/internal/controller/note/delete"
	"cloudos/internal/controller/note/deltag"
	"cloudos/internal/controller/note/info"
	"cloudos/internal/controller/note/list"
	"cloudos/internal/controller/note/save"
	"cloudos/internal/controller/note/topics"

	"github.com/gin-gonic/gin"
)

// NoteList
// @summary 文档列表
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/v1/note/list [post]
func (a *Api) NoteList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}

// NoteSave
// @summary 保存笔记
// @Tags 笔记
// @Accept json
// @Produce json
// @Param param body save.Params true "请求参数"
// @Response 200 object system.Response{data=save.Reply} "调用成功"
// @Router /api/v1/note/save [post]
func (a *Api) NoteSave(ctx *gin.Context) {
	a.Scheduler(save.NewController(ctx))
}

// NoteDelete
// @summary 删除文档
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/v1/note/delete [post]
func (a *Api) NoteDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// NoteAddtag
// @summary 添加标签
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body addtag.Params true "请求参数"
// @Response 200 object system.Response{data=addtag.Reply} "调用成功"
// @Router /api/v1/note/addtag [post]
func (a *Api) NoteAddtag(ctx *gin.Context) {
	a.Scheduler(addtag.NewController(ctx))
}

// NoteDeltag
// @summary 删除标签
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body deltag.Params true "请求参数"
// @Response 200 object system.Response{data=deltag.Reply} "调用成功"
// @Router /api/v1/note/deltag [post]
func (a *Api) NoteDeltag(ctx *gin.Context) {
	a.Scheduler(deltag.NewController(ctx))
}

// NoteTopics
// @summary 主题列表
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body topics.Params true "请求参数"
// @Response 200 object system.Response{data=topics.Reply} "调用成功"
// @Router /api/v1/note/topics [get]
func (a *Api) NoteTopics(ctx *gin.Context) {
	a.Scheduler(topics.NewController(ctx))
}

// NoteInfo
// @summary 笔记详情
// @Tags 文档
// @Accept json
// @Produce json
// @Param param body info.Params true "请求参数"
// @Response 200 object system.Response{data=info.Reply} "调用成功"
// @Router /api/v1/note/info [get]
func (a *Api) NoteInfo(ctx *gin.Context) {
	a.Scheduler(info.NewController(ctx))
}
