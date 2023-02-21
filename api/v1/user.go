package v1

import (
	"superserver/internal/controller/user/login"

	"github.com/gin-gonic/gin"
)

// UserLogin
// @summary 用户登录
// @Accept json
// @Produce json
// @Param param body login.Param true "请求参数"
// @Response 200 object login.Reply "调用成功"
// @Response 400 object system.ChainContext "请求错误"
// @Response 500 object system.ChainContext "系统错误"
// @Router /api/v1/user/login [post]
func UserLogin(ctx *gin.Context) {
	if controller, err := login.NewController(ctx); err == nil {
		controller.Deal()
	}
}
