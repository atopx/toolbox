package v1

import (
	"toolbox/internal/controller/home/wol"

	"github.com/gin-gonic/gin"
)

// HomeWol
// @summary 远程开机
// @Accept json
// @Produce json
// @Param param query wol.Param true "请求参数"
// @Response 200 object wol.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/home/wol [get]
func HomeWol(ctx *gin.Context) {
	if controller, err := wol.NewController(ctx); err == nil {
		controller.Deal()
	}
}
