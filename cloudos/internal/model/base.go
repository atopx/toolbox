package model

import (
	"cloudos/common/system"
	"fmt"

	"gorm.io/gorm"
)

type Base struct{}

func (b *Base) Db() *gorm.DB {
	return system.GetHandler().Db
}

func (dao *Base) NotDeleted(tx *gorm.DB) *gorm.DB {
	return tx.Where("delete_time = 0")
}

func (dao *Base) Like(value string) string {
	return fmt.Sprintf("%%%s%%", value)
}
