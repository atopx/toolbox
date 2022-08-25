package v1

import (
	"toolbox/internal/controller/json/format"
	"toolbox/internal/controller/json/trans"

	"github.com/gin-gonic/gin"
)

// JsonTrans
// @summary json去除转义
// @Accept json
// @Produce json
// @Param param query trans.Param true "请求参数"
// @Response 200 object trans.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/json/trans [get]
func JsonTrans(ctx *gin.Context) {
	if controller, err := trans.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// JsonFormat
// @summary json格式化
// @Accept json
// @Produce json
// @Param param query format.Param true "请求参数"
// @Response 200 object format.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/json/format [get]
func JsonFormat(ctx *gin.Context) {
	if controller, err := format.NewController(ctx); err == nil {
		controller.Deal()
	}
}
