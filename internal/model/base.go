package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id         int64 `json:"id"`
	DeleteTime int64 `json:"delete_time"`
	CreateTime int64 `json:"create_ime"`
	UpdateTime int64 `json:"update_ime"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	timestamp := time.Now().Local().Unix()
	b.CreateTime = timestamp
	b.UpdateTime = timestamp
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	timestamp := time.Now().Local().Unix()
	b.UpdateTime = timestamp
	return nil
}

type BaseDao struct {
	Db        *gorm.DB
	Tablename string
}

func (dao *BaseDao) Connection() *gorm.DB {
	return dao.Db.Table(dao.Tablename)
}

func (dao *BaseDao) Load(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).First(po).Error
}

func (dao *BaseDao) Create(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).Create(po).Error
}

func (dao *BaseDao) Update(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).Updates(po).Error
}

func (dao *BaseDao) Save(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).Save(po).Error
}

func (dao *BaseDao) Delete(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).Update("delete_time", time.Now().Local().Unix()).Error
}

func (dao *BaseDao) RealDelete(po Po) error {
	return dao.Connection().Where("id = ?", po.GetPk()).Delete(po).Error
}

type Po interface {
	GetPk() any
}

type Enum struct {
	Enum  string `json:"enum"`
	Value string `json:"value"`
}
