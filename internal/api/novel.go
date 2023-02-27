package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/novel/create"
)

// NovelCreate
// @summary 创建小说
// @Tags 角色
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
