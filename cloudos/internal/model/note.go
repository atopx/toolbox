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

func (dao *NoteDao) CreateNote(note *pb.Note) error {
	note.CreateTime = time.Now().Local().Unix()
	note.UpdateTime = note.CreateTime
	return dao.Db().Create(note).Error
}

func (dao *NoteDao) UpdateNote(note *pb.Note) error {
	if note.GetId() == 0 {
		return errors.New("missing note.id")
	}
	note.UpdateTime = time.Now().Local().Unix()
	return dao.Db().Updates(note).Error
}

func (dao *NoteDao) DeleteNote(note *pb.Note) error {
	if note.GetId() == 0 {
		return errors.New("missing note.id")
	}
	ts := time.Now().Local().Unix()
	return dao.Db().Model(note).Where("id = ?", note.Id).UpdateColumn("delete_time", ts).Error
}

func (dao *NoteDao) First(query any, args ...any) *pb.Note {
	note := new(pb.Note)
	err := dao.Db().Scopes(dao.NotDeleted).Where(query, args...).First(&note).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return note
}
