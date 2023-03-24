package login

import (
	"errors"
	"superserver/common/interface/ecode_iface"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/internal/model/user"
	"superserver/internal/model/user_token"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	if params.Username == "" || params.Password == "" {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "用户名或密码错误")
		return
	}
	dao := user.NewDao(ctl.GetDatabase())

	po, err := dao.GetUserByUsername(params.Username, true)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "用户不存在")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "user login dao.GetUserByUsername failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12001]系统错误, 请联系管理员")
		return
	}
	if !po.ComparePassword(params.Password) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "登录密码错误")
		return
	}
	// 如果已有token，直接返回给用户
	tokenDao := user_token.NewDao(ctl.GetDatabase())
	token, err := tokenDao.First(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", po.Id)
	})

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error(ctl.Context, "user login tokenDao.First failed", zap.Error(err))
			ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12001]系统错误, 请联系管理员")
			return
		}
	}
	// 如果没有token, 或者token过期， 需要重新签发
	if token.ExpireTime <= time.Now().Unix() {
		newToken := system.SignClaims(po.Id)
		token.AccessToken = newToken.AccessToken
		token.RefreshToken = newToken.RefreshToken
		token.ExpireTime = newToken.IssuedTime
		token.UserId = po.Id
	}

	if err = tokenDao.UpsertByUserId(token); err != nil {
		logger.Error(ctl.Context, "user login system.SignClaims failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12002]系统错误, 请联系管理员")
		return
	}

	ctl.NewOkResponse(&Reply{
		UserId:       po.Id,
		Name:         po.Name,
		Username:     po.Username,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expires:      token.ExpireTime,
	})
}
