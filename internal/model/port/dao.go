package port

import (
	"superserver/internal/model"

	"gorm.io/gorm"
)

type PortDao struct {
	model.BaseDao
}

func (*PortDao) Tablename() string {
	return "su_computer_port"
}

func NewDao(db *gorm.DB) *PortDao {
	dao := &PortDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.Tablename = dao.Tablename()
	once.Do(func() { dao.init() })
	return dao
}

func (dao *PortDao) init() {
	if !dao.Db.Migrator().HasTable(dao.Tablename) {
		err := dao.Db.Exec(`CREATE TABLE IF NOT EXISTS su_computer_port (
			id integer PRIMARY KEY AUTOINCREMENT,
			computer_id integer not null, -- 主机ID
			port integer not null default 0, -- 网络端口 
			protocol tinyint not null default 0, -- 应用协议
			tranport tinyint not null default 0, -- 传输协议
			use_type tinyint not null default 0, -- 用途
			desc text not null default '', -- 备注
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
