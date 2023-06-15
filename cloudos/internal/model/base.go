package model

import (
	"cloudos/common/system"

	"gorm.io/gorm"
)

type Base struct{}

func (b *Base) Db() *gorm.DB {
	return system.GetHandler().Db
}

func (dao *Base) NotDeleted(tx *gorm.DB) *gorm.DB {
	return tx.Where("delete_time = 0")
}
