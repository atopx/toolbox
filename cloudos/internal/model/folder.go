package model

import (
	"cloudos/common/pb"
	"errors"
	"time"

	"gorm.io/gorm"
)

type FolderDao struct {
	Base
}

func (dao *FolderDao) Create(obj *pb.Folder) error {
	obj.CreateTime = time.Now().Local().Unix()
	obj.UpdateTime = obj.CreateTime
	return dao.Db().Create(obj).Error
}

func (dao *FolderDao) Update(obj *pb.Folder) error {
	obj.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(obj).Error
}

func (dao *FolderDao) Delete(obj *pb.Folder) error {
	ts := time.Now().Local().Unix()
	return dao.Db().Model(obj).Where("id = ?", obj.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *FolderDao) First(query any, args ...any) *pb.Folder {
	obj := new(pb.Folder)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&obj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return obj
}

func (dao *FolderDao) Pager(tx *gorm.DB, pager *pb.Pager) (objs []*pb.Folder) {
	tx.Offset(int(pager.Size * (pager.Index - 1))).Limit(int(pager.Size)).
		Order("update_time desc").Find(&objs)
	return objs
}
