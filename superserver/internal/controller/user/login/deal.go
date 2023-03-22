package login

import (
	"errors"
	"net/http"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/internal/model/user"
	"superserver/internal/model/user_token"

	"go.uber.org/zap"
	"gorm.io/gorm"
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
	token := system.SignClaims(po.Id)
	token.UserId = po.Id
	tokenDao := user_token.NewDao(ctl.GetDatabase())
	if err = tokenDao.Create(&token); err != nil {
		logger.Error(ctl.Context, "user login system.SignClaims failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "[12002]系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(http.StatusOK, &Reply{
		UserId:       po.Id,
		Name:         po.Name,
		Username:     po.Username,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expires:      token.ExpireTime,
	})
}
