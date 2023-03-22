package refresh

import (
	"errors"
	"superserver/common/interface/ecode_iface"
	"superserver/common/logger"
	"superserver/common/system"
	"superserver/internal/model/user_token"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)

	tokenDao := user_token.NewDao(ctl.GetDatabase())
	token, err := tokenDao.First(func(db *gorm.DB) *gorm.DB {
		return db.Where("refresh_token = ?", params.RefreshToken)
	})
	if errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_BAD_REQUEST, "无效的TOKEN")
		return
	}
	if err != nil {
		logger.Error(ctl.Context, "user refresh token tokenDao.First failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12001]系统错误, 请联系管理员")
		return
	}

	newToken := system.SignClaims(token.UserId)
	newToken.Id = token.Id
	err = tokenDao.Update(&newToken)
	if err != nil {
		logger.Error(ctl.Context, "user refresh token tokenDao.Update failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "[12001]系统错误, 请联系管理员")
		return
	}
	ctl.NewOkResponse(&Reply{
		AccessToken:  newToken.AccessToken,
		RefreshToken: newToken.RefreshToken,
		Expires:      newToken.ExpireTime,
	})
}
