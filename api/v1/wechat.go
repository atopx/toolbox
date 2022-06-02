package v1

import (
	"toolbox/internal/controller/wechat/auth_token"
	"toolbox/internal/controller/wechat/message_listen"

	"github.com/gin-gonic/gin"
)

// WechatAuthToken
// @summary 微信公众号token认证
// @Accept json
// @Produce json
// @Param param query spec.Param true "请求参数"
// @Response 200 object spec.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/wechat [get]
func WechatAuthToken(ctx *gin.Context) {
	if controller, err := auth_token.NewController(ctx); err == nil {
		controller.Deal()
	}
}

// WechatMessageListen
// @summary 微信公众号消息监听
// @Accept xml
// @Produce json
// @Param param body message_listen.Param true "请求参数"
// @Response 200 object message_listen.Reply "调用成功"
// @Response 400 object system.ReplyError "请求错误"
// @Response 500 object system.ReplyError "系统错误"
// @Router /api/v1/wechat [post]
func WechatMessageListen(ctx *gin.Context) {
	if controller, err := message_listen.NewController(ctx); err == nil {
		controller.Deal()
	}
}
