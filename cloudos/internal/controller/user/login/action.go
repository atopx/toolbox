package login

import (
	"cloudos/common/logger"
	"cloudos/common/pb"
	"cloudos/common/utils"
	"cloudos/internal/model"
	"time"

	"go.uber.org/zap"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)
	if params.Username == "" {
		return nil, pb.ECode_InvalidUsername
	}
	if params.Password == "" {
		return nil, pb.ECode_InvalidPassword
	}

	user := new(model.UserDao).First("username = ?", params.Username)
	if user == nil {
		return nil, pb.ECode_NotFoundUser
	}

	if utils.Hash(params.Password) != user.Password {
		return nil, pb.ECode_InvalidPassword
	}

	authTokenDao := new(model.AuthTokenDao)
	token := authTokenDao.First("user_id = ?", user.Id)

	current := time.Now().Local()

	if token == nil {
		token = new(pb.AuthToken)
	}

	if token.ExpireTime <= current.Unix() {
		// 没有token或已过期， 需要重新签发
		expires := current.Add(24 * time.Hour)
		token.UserId = user.Id
		token.IssuedTime = current.Unix()
		token.ExpireTime = expires.Unix()
		token.AccessToken = utils.SignToken(current, expires, token.ExpireTime)
		// 使用AccessToken的过期时间加密
		token.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), token.ExpireTime)
	}

	// save token
	if err := authTokenDao.Save(token); err != nil {
		logger.Error(c.Context(), "save token error", zap.Error(err))
		return token, pb.ECode_ServerInternalError
	}

	return token, pb.ECode_SUCCESS
}
