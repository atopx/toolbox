package models

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	conn *gorm.DB `gorm:"-"`

	Id         int   `json:"id"`                                // 主键
	CreateTime int64 `json:"create_time" gorm:"autoCreateTime"` // 创建时间 时间戳：秒
	UpdateTime int64 `json:"update_time" gorm:"autoUpdateTime"` // 更新时间 时间戳：秒
	DeleteTime int64 `json:"delete_time"`                       // 删除时间 时间戳：秒
}

type Pager struct {
	Index int   `json:"index"`
	Size  int   `json:"size"`
	Total int64 `json:"total"`
}

func (b *Base) Connect() *gorm.DB {
	return b.conn.Table("base")
}

func (b *Base) BeforeCreate(*gorm.DB) error {
	timestamp := time.Now().Local().Unix()
	b.CreateTime = timestamp
	b.UpdateTime = timestamp
	return nil
}

func (b *Base) BeforeUpdate(*gorm.DB) error {
	timestamp := time.Now().Local().Unix()
	b.UpdateTime = timestamp
	return nil
}

func (b *Base) Create(obj Object) error {
	return obj.Connect().Create(obj).Error
}

func (b *Base) Update(obj Object) error {
	return obj.Connect().Updates(obj).Error
}

func (b *Base) Save(obj Object) error {
	return obj.Connect().Save(obj).Error
}

func (b *Base) NotDeleted(tx *gorm.DB) *gorm.DB {
	return tx.Where("delete_time = 0")
}

func (b *Base) Paginate(pager *Pager) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset((pager.Index - 1) * pager.Size).Limit(pager.Size)
	}
}

type Object interface {
	Connect() *gorm.DB
}

const (
	TablenameBook    = "book"
	TablenameChapter = "chapter"
)
