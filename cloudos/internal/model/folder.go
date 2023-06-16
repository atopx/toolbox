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

func (dao *FolderDao) CreateFolder(folder *pb.Folder) error {
	folder.CreateTime = time.Now().Local().Unix()
	folder.UpdateTime = folder.CreateTime
	return dao.Db().Create(folder).Error
}

func (dao *FolderDao) UpdateFolder(folder *pb.Folder) error {
	if folder.GetId() == 0 {
		return errors.New("missing folder.id")
	}
	folder.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(folder).Error
}

func (dao *FolderDao) DeleteFolder(folder *pb.Folder) error {
	if folder.GetId() == 0 {
		return errors.New("missing folder.id")
	}
	ts := time.Now().Local().Unix()
	return dao.Db().Model(folder).Where("id = ?", folder.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *FolderDao) First(query any, args ...any) *pb.Folder {
	folder := new(pb.Folder)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&folder).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return folder
}

func (dao *FolderDao) Pager(tx *gorm.DB, pager *pb.Pager) (folders []*pb.Folder) {
	tx.Offset(int(pager.Size * (pager.Index - 1))).Limit(int(pager.Size)).
		Order("update_time desc").Find(&folders)
	return folders
}
