package model

import (
	"cloudos/common/pb"
	"cloudos/common/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserDao struct {
	Base
}

func (dao *UserDao) CreateUser(user *pb.User) error {
	user.Password = utils.Hash(user.Password)
	user.CreateTime = time.Now().Local().Unix()
	user.UpdateTime = user.CreateTime
	return dao.Db().Create(user).Error
}

func (dao *UserDao) UpdateUser(user *pb.User) error {
	if user.GetId() == 0 {
		return errors.New("missing user.id")
	}
	user.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(user).Error
}

func (dao *UserDao) DeleteUser(user *pb.User) error {
	if user.GetId() == 0 {
		return errors.New("missing user.id")
	}
	ts := time.Now().Local().Unix()
	return dao.Db().Model(user).Where("id = ?", user.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *UserDao) First(query any, args ...any) *pb.User {
	user := new(pb.User)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}
