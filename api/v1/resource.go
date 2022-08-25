package v1

import (
	"toolbox/internal/controller/resource"

	"github.com/gin-gonic/gin"
)

// ResourceGet
// @summary 获取文件资源
// @Accept json
// @Produce json
// @Param param query resource.Param true "请求参数"
// @Response 200 object string "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/resource/get [get]
func ResourceGet(ctx *gin.Context) {
	if controller, err := resource.NewController(ctx); err == nil {
		controller.Deal()
	}
}
