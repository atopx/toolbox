package login

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/internal/model/user"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Username == "" || params.Password == "" {
		ctl.NewErrorResponse(http.StatusBadRequest, "用户名或密码错误")
		return
	}
	dao := user.NewDao(ctl.GetDatabase())

	po, err := dao.GetUserByUsername(params.Username, true)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(http.StatusBadRequest, "用户不存在")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "user login dao.GetUserByUsername failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "[12001]系统错误, 请联系管理员")
		return
	}
	if !po.ComparePassword(params.Password) {
		ctl.NewErrorResponse(http.StatusBadRequest, "登录密码错误")
		return
	}
	tokenStr, err := system.SignClaims(po.Id, po.RoleId)
	if err != nil {
		logger.Error(ctl.Context, "user login system.SignClaims failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "[12002]系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(http.StatusOK, &Reply{UserId: po.Id, Token: tokenStr})
}
