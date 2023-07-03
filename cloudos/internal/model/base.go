package model

import (
	"cloudos/common/pb"
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

// Paginate 分页scope
func (dao *Base) Paginate(pager *pb.Pager) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pager.Disabled {
			return db
		}
		offset := (pager.Index - 1) * pager.Size
		return db.Offset(int(offset)).Limit(int(pager.Size))
	}
}
