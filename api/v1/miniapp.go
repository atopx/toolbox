package v1

import (
	"toolbox/internal/controller/wechat/miniapp"

	"github.com/gin-gonic/gin"
)

// Scheduler
// @summary 命令调度器
// @Accept json
// @Produce json
// @Param param query miniapp.Param true "请求参数"
// @Response 200 object miniapp.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/wechat/scheduler [get]
func MiniappScheduler(ctx *gin.Context) {
	if controller, err := miniapp.NewController(ctx); err == nil {
		controller.Deal()
	}
}
