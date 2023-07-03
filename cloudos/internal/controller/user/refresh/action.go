package refresh

import (
	"cloudos/common/consts"
	"cloudos/common/logger"
	"cloudos/common/pb"
	"cloudos/common/utils"
	"cloudos/internal/model"
	"time"

	"go.uber.org/zap"
)

func (c *Controller) Deal() (any, pb.ECode) {
	params := c.Params.(*Params)
	if params.RefreshToken == consts.EmptyStr {
		return nil, pb.ECode_AuthTokenInvalid
	}

	authTokenDao := new(model.AuthTokenDao)
	token := authTokenDao.First("refresh_token = ?", params.RefreshToken)

	current := time.Now().Local()

	if token == nil {
		return nil, pb.ECode_AuthTokenInvalid
	}

	// 重新签发token
	expires := current.Add(24 * time.Hour)
	token.IssuedTime = current.Unix()
	token.ExpireTime = expires.Unix()
	token.AccessToken = utils.SignToken(current, expires, token.ExpireTime)
	// 使用AccessToken的过期时间加密
	token.RefreshToken = utils.SignToken(current, current.Add(7*24*time.Hour), token.ExpireTime)

	// save token
	if err := authTokenDao.Save(token); err != nil {
		logger.Error(c.Context(), "save token error", zap.Error(err))
		return nil, pb.ECode_ServerInternalError
	}

	return c.buildReply(token), pb.ECode_SUCCESS
}

func (c *Controller) buildReply(token *pb.AuthToken) *Reply {
	return &Reply{
		UserId:       token.UserId,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpireTime:   token.ExpireTime,
	}
}
