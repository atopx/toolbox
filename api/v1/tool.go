package v1

import (
	"toolbox/internal/controller/tool/dblog"

	"github.com/gin-gonic/gin"
)

// DbLogDraw
// @summary DB日志解析SQL
// @Accept json
// @Produce json
// @Param param query dblog.Param true "请求参数"
// @Response 200 object dblog.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/tool/dblog [get]
func DbLogDraw(ctx *gin.Context) {
	if controller, err := dblog.NewController(ctx); err == nil {
		controller.Deal()
	}
}
