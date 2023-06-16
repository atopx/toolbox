package model

import (
	"cloudos/common/pb"
	"errors"
	"time"

	"gorm.io/gorm"
)

type AuthTokenDao struct {
	Base
}

func (dao *AuthTokenDao) CreateAuthToken(authToken *pb.AuthToken) error {
	authToken.CreateTime = time.Now().Local().Unix()
	authToken.UpdateTime = authToken.CreateTime
	return dao.Db().Create(authToken).Error
}

func (dao *AuthTokenDao) UpdateAuthToken(authToken *pb.AuthToken) error {
	if authToken.GetId() == 0 {
		return errors.New("missing authToken.id")
	}
	authToken.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(authToken).Error
}

func (dao *AuthTokenDao) DeleteAuthToken(authToken *pb.AuthToken) error {
	if authToken.GetId() == 0 {
		return errors.New("missing authToken.id")
	}
	ts := time.Now().Local().Unix()
	return dao.Db().Model(authToken).Where("id = ?", authToken.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *AuthTokenDao) First(query any, args ...any) *pb.AuthToken {
	authToken := new(pb.AuthToken)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&authToken).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return authToken
}

func (dao *AuthTokenDao) Save(token *pb.AuthToken) error {
	ts := time.Now().Unix()
	switch token.Id {
	case 0:
		token.CreateTime = ts
		token.UpdateTime = ts
	default:
		token.UpdateTime = ts
	}
	return dao.Db().Save(token).Error
}
