package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/label/create"
	"superserver/internal/controller/label/delete"
	"superserver/internal/controller/label/list"
)

// LabelCreate
// @summary 创建标签
// @Tags 标签
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response{data=create.Reply} "调用成功"
// @Router /api/label/create [post]
func (a *Api) LabelCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// LabelDelete
// @summary 删除标签
// @Tags 标签
// @Accept json
// @Produce json
// @Param param body delete.Params true "请求参数"
// @Response 200 object system.Response{data=delete.Reply} "调用成功"
// @Router /api/label/delete [delete]
func (a *Api) LabelDelete(ctx *gin.Context) {
	a.Scheduler(delete.NewController(ctx))
}

// LabelList
// @summary 标签列表
// @Tags 标签
// @Accept json
// @Produce json
// @Param param body list.Params true "请求参数"
// @Response 200 object system.Response{data=list.Reply} "调用成功"
// @Router /api/label/list [post]
func (a *Api) LabelList(ctx *gin.Context) {
	a.Scheduler(list.NewController(ctx))
}
