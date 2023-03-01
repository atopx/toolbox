package api

import (
	"superserver/internal/controller/novel/create"
	"superserver/internal/controller/novel/export"

	"github.com/gin-gonic/gin"
)

// NovelCreate
// @summary 创建小说
// @Tags 小说
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/novel/create [post]
func (a *Api) NovelCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// NovelExport
// @summary 创建小说
// @Tags 小说
// @Accept json
// @Produce json
// @Param param body export.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/novel/export [get]
func (a *Api) NovelExport(ctx *gin.Context) {
	a.Scheduler(export.NewController(ctx))
}
