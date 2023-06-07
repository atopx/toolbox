package api

import (
	"cloudos/internal/controller/user/create"
	"cloudos/internal/controller/user/info"
	"cloudos/internal/controller/user/login"
	"cloudos/internal/controller/user/refresh"

	"github.com/gin-gonic/gin"
)

// UserCreate
// @summary 用户注册
// @Tags 用户
// @Accept json
// @Produce json
// @Param param body create.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
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
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
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
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
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
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/user/info [post]
func (a *Api) UserInfo(ctx *gin.Context) {
	a.Scheduler(info.NewController(ctx))
}
