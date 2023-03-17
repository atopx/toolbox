package model

import (
	"fmt"
	"superserver/common/utils"
	"superserver/proto/common_iface"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Id         int   `json:"id"`                                // 自增ID
	DeleteTime int64 `json:"delete_time"`                       // 更新时间
	CreateTime int64 `json:"create_time" gorm:"autoCreateTime"` // 创建时间
	UpdateTime int64 `json:"update_time" gorm:"autoUpdateTime"` // 删除时间
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

type BaseDao struct {
	Db        *gorm.DB
	TableName string
}

func (dao *BaseDao) Connection() *gorm.DB {
	return dao.Db.Table(dao.TableName)
}

func (dao *BaseDao) Load(po Po) error {
	return dao.Connection().First(po).Error
}

func (dao *BaseDao) Create(po Po) error {
	return dao.Connection().Create(po).Error
}

func (dao *BaseDao) Update(po Po) error {
	return dao.Connection().Where("id = ?", po.GetId()).Updates(po).Error
}

func (dao *BaseDao) Save(po Po) error {
	return dao.Connection().Where("id = ?", po.GetId()).Save(po).Error
}

func (dao *BaseDao) Delete(po Po) error {
	return dao.Connection().Where("id = ?", po.GetId()).Update("delete_time", time.Now().Local().Unix()).Error
}

func (dao *BaseDao) RealDelete(po Po) error {
	return dao.Connection().Where("id = ?", po.GetId()).Delete(po).Error
}

func (dao *BaseDao) NotDeleted(tx *gorm.DB) *gorm.DB {
	return tx.Where("delete_time = 0")
}

func (dao *BaseDao) Range(key string, rang *common_iface.RangeI64) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if rang == nil {
			return tx
		}
		left, right := utils.NewRangeValue(rang)
		return tx.Where(fmt.Sprintf("%s %s %d and %s %s %d", key, left, rang.Left, key, right, rang.Right))
	}
}

func (dao *BaseDao) Paginate(pager *common_iface.Pager) func(tx *gorm.DB) *gorm.DB {
	return func(tx *gorm.DB) *gorm.DB {
		if pager == nil || pager.Disabled {
			return tx
		}
		offset := (pager.Index - 1) * pager.Size
		return tx.Offset(int(offset)).Limit(int(pager.Size))
	}
}

type Po interface {
	GetId() any
}

type Enum struct {
	Enum  string `json:"enum"`
	Value string `json:"value"`
}
