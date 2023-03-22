package user_token

import "superserver/internal/model"

type UserToken struct {
	model.Base
	UserId       int
	AccessToken  string
	RefreshToken string
	ExpireTime   int64
	IssuedTime   int64
}
