package model

import (
	"cloudos/common/pb"
	"errors"
	"time"

	"gorm.io/gorm"
)

type NoteDao struct {
	Base
}

func (dao *NoteDao) Create(obj *pb.Note) error {
	obj.CreateTime = time.Now().Local().Unix()
	obj.UpdateTime = obj.CreateTime
	return dao.Db().Create(obj).Error
}

func (dao *NoteDao) Update(obj *pb.Note) error {
	obj.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(obj).Error
}

func (dao *NoteDao) Delete(obj *pb.Note) error {
	ts := time.Now().Local().Unix()
	return dao.Db().Model(obj).Where("id = ?", obj.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *NoteDao) First(query any, args ...any) *pb.Note {
	obj := new(pb.Note)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return obj
}
