package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/user/create"
	"superserver/internal/controller/user/info"
	"superserver/internal/controller/user/login"
	"superserver/internal/controller/user/refresh"
)

// UserCreate
// @summary 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.Response "调用成功"
// @Router /api/user/create [post]
func (a *Api) UserCreate(ctx *gin.Context) {
	a.Scheduler(create.NewController(ctx))
}

// UserLogin
// @summary 用户登录
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body login.Params true "请求参数"
// @Response 200 object system.Response "调用成功"
// @Router /api/user/login [post]
func (a *Api) UserLogin(ctx *gin.Context) {
	a.Scheduler(login.NewController(ctx))
}

// UserRefresh
// @summary 用户刷新AccessToken
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body refresh.Params true "请求参数"
// @Response 200 object system.Response "调用成功"
// @Router /api/user/refresh [post]
func (a *Api) UserRefresh(ctx *gin.Context) {
	a.Scheduler(refresh.NewController(ctx))
}

// UserInfo
// @summary 用户详情
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body info.Params true "请求参数"
// @Response 200 object system.Response "调用成功"
// @Router /api/user/info [get]
func (a *Api) UserInfo(ctx *gin.Context) {
	a.Scheduler(info.NewController(ctx))
}
