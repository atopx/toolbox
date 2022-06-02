package v1

import (
	"toolbox/internal/controller/pixel/spec"

	"github.com/gin-gonic/gin"
)

// PixelSpec
// @summary 计算显示器实际规格
// @Accept json
// @Produce json
// @Param param query spec.Param true "请求参数"
// @Response 200 object spec.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/pixel/spec [get]
func PixelSpec(ctx *gin.Context) {
	if controller, err := spec.NewController(ctx); err == nil {
		controller.Deal()
	}
}
