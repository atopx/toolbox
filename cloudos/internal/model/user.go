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

func (dao *UserDao) CreateUser(obj *pb.User) error {
	obj.Password = utils.Hash(obj.Password)
	obj.CreateTime = time.Now().Local().Unix()
	obj.UpdateTime = obj.CreateTime
	return dao.Db().Create(obj).Error
}

func (dao *UserDao) Update(obj *pb.User) error {
	obj.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(obj).Error
}

func (dao *UserDao) Delete(obj *pb.User) error {
	ts := time.Now().Local().Unix()
	return dao.Db().Model(obj).Where("id = ?", obj.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *UserDao) First(query any, args ...any) *pb.User {
	obj := new(pb.User)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return obj
}
