package computer

import (
	"fmt"
	"superserver/common/utils"
	"superserver/internal/model"
	"superserver/proto/common_iface"

	"gorm.io/gorm"
)

type ComputerDao struct {
	model.BaseDao
}

func (*ComputerDao) Tablename() string {
	return "su_computer"
}

func NewDao(db *gorm.DB) *ComputerDao {
	dao := &ComputerDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.Tablename = dao.Tablename()
	once.Do(func() { dao.init() })
	return dao
}

func (dao *ComputerDao) FilterBy(key string, value any) (po *Computer, err error) {
	err = dao.Connection().Where(fmt.Sprintf("%s = ?", key), value).First(&po).Error
	return po, err
}

func (dao *ComputerDao) Filter(filter *Filter, pager *common_iface.Pager) (computers []Computer, err error) {
	if filter == nil {
		filter = new(Filter)
	}
	if pager == nil {
		pager = &common_iface.Pager{Disabled: true}
	}
	tx := dao.Connection().Where("delete_time = 0")
	if filter.Keyword != "" {
		tx.Where("name like ? or username like ?", utils.NewLikeValue(filter.Keyword), utils.NewLikeValue(filter.Keyword))
	}

	if filter.Hostname != "" {
		tx.Where("lan_ostname like ? or wan_hostname like ?", utils.NewLikeValue(filter.Hostname), utils.NewLikeValue(filter.Hostname))
	}

	if filter.Address != "" {
		filter.Address = utils.MACEncode(filter.Address)
		tx.Where("address like ?", utils.NewLikeValue(filter.Address))
	}

	if len(filter.PowerStatus) > 0 {
		tx.Where("power_status in ?", filter.PowerStatus)
	}
	tx.Scopes(dao.Range("create_time", filter.CreateTimeRange))
	tx.Scopes(dao.Range("update_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Range("scan_time", filter.UpdateTimeRange))
	tx.Scopes(dao.Paginate(pager))
	tx.Count(&pager.Count)
	err = tx.Find(&computers).Error
	return computers, err
}

func (dao *ComputerDao) init() {
	if !dao.Db.Migrator().HasTable(dao.Tablename()) {
		err := dao.Db.Exec(`CREATE TABLE IF NOT EXISTS su_computer (
			id integer PRIMARY KEY AUTOINCREMENT, -- 自增ID
			name varchar ( 64 )  NOT NULL, -- 名称
			username varchar ( 64 ) NOT NULL DEFAULT '', -- 用户名
			password varchar ( 64 ) NOT NULL DEFAULT '', -- 用户密码
			lan_hostname varchar ( 64 ) NOT NULL DEFAULT '', -- 局域网地址
			wan_hostname varchar ( 64 ) NOT NULL DEFAULT '', -- 广域网地址
			address char ( 12 ) NOT NULL DEFAULT '', -- 物理地址
			power_status tinyint NOT NULL DEFAULT 0, -- 电源状态
			scan_time bigint NOT NULL DEFAULT 0, -- 最后一次扫描时间
			creator integer not null, -- 创建人
			updator integer not null, -- 更新人
			create_time bigint not null, -- 创建时间
			update_time bigint not null, -- 更新时间
			delete_time bigint not null default 0 -- 删除时间
		);`).Error
		if err != nil {
			panic(err)
		}
	}

}
