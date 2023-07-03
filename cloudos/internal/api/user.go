package api

import (
	"cloudos/internal/controller/user/login"
	"cloudos/internal/controller/user/refresh"

	"github.com/gin-gonic/gin"
)

// UserLogin
// @summary 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body login.Params true "请求参数"
// @Response 200 object system.Response{data=login.Reply} "调用成功"
// @Router /api/v1/user/login [post]
func (a *Api) UserLogin(ctx *gin.Context) {
	a.Scheduler(login.NewController(ctx))
}

// UserRefresh
// @summary 刷新token
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body refresh.Params true "请求参数"
// @Response 200 object system.Response{data=refresh.Reply} "调用成功"
// @Router /api/v1/user/refresh [post]
func (a *Api) UserRefresh(ctx *gin.Context) {
	a.Scheduler(refresh.NewController(ctx))
}
