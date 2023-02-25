package api

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller/user/create"
)

// UserCreate
// @summary 用户注册、创建用户
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
// @Accept json
// @Produce json
// @Param param body login.Params true "请求参数"
// @Response 200 object system.ChainMessage "调用成功"
// @Response 400 object system.ChainMessage "请求错误"
// @Response 500 object system.ChainMessage "系统错误"
// @Router /api/user/create [post]
func (a *Api) UserLogin(ctx *gin.Context) {
	//a.Scheduler(login.NewController(ctx))
}
