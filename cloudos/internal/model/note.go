package model

import (
	"cloudos/common/pb"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NoteDao struct {
	Base
}

func (dao *NoteDao) Save(obj *pb.Note) error {
	ts := time.Now().Unix()
	switch obj.Id {
	case 0:
		obj.CreateTime = ts
		obj.UpdateTime = ts
	default:
		obj.UpdateTime = ts
	}
	return dao.Db().Save(obj).Error
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

func (dao *NoteDao) Labels(noteId int64) []string {
	labels := make([]string, 0)
	dao.Db().Model(new(pb.NoteLabel)).
		Scopes(dao.NotDeleted).
		Where("note_id = ?", noteId).
		Select("name").
		Scan(&labels)
	return labels
}

func (dao *NoteDao) CleanLabels(noteId int64) error {
	return dao.Db().Model(new(pb.NoteLabel)).Delete("note_id = ?", noteId).Error
}

func (dao *NoteDao) AddLabels(noteId int64, labels []string) error {
	if num := len(labels); num > 0 {
		labelObjs := make([]*pb.NoteLabel, 0, num)
		ts := time.Now().Unix()
		for i := 0; i < num; i++ {
			labelObjs = append(labelObjs, &pb.NoteLabel{
				Name:       labels[i],
				NoteId:     noteId,
				CreateTime: ts,
				UpdateTime: ts,
			})
		}
		return dao.Db().Model(new(pb.NoteLabel)).Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "note_id"}, {Name: "name"}},
			DoNothing: true,
		}).Create(&labelObjs).Error
	}
	return nil
}

func (dao *NoteDao) NoteLabels(notes []*pb.Note) map[int64][]string {
	result := make(map[int64][]string)
	if len(notes) == 0 {
		return result
	}
	var noteIds []int64
	for _, note := range notes {
		noteIds = append(noteIds, note.Id)
	}
	var labels []*pb.NoteLabel
	dao.Db().Scopes(dao.NotDeleted).Where("note_id in ?", noteIds).Find(&labels)
	for _, label := range labels {
		result[label.NoteId] = append(result[label.NoteId], label.Name)
	}
	return result
}

func (dao *NoteDao) ListSelectFields() []string {
	return []string{"title", "topic", "create_time", "update_time"}
}
