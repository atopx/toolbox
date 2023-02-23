package user

import (
	"superserver/internal/model"
	"superserver/proto/user_iface"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UserDao struct {
	model.BaseDao
}

func (*UserDao) Tablename() string {
	return "su_user"
}

func NewDao(db *gorm.DB) *UserDao {
	dao := &UserDao{BaseDao: model.BaseDao{Db: db}}
	dao.BaseDao.Tablename = dao.Tablename()
	once.Do(func() { dao.init() })
	return dao
}

func (dao *UserDao) init() {
	if !dao.Db.Migrator().HasTable(dao.Tablename) {
		err := dao.Db.Exec(`CREATE TABLE IF NOT EXISTS su_user (
			id integer PRIMARY KEY AUTOINCREMENT, -- 自增ID
			name varchar ( 64 )  NOT NULL, -- 名称
			username varchar ( 64 ) NOT NULL DEFAULT '', -- 用户名
			password varchar ( 64 ) NOT NULL DEFAULT '', -- 用户密码
			role tinyint not null default 0, -- 用户角色
			status tinyint not null default 0, -- 用户状态
			create_time bigint not null, -- 创建时间
			update_time bigint not null, -- 更新时间
			delete_time bigint not null default 0 -- 删除时间
		);`).Error
		if err != nil {
			panic(err)
		}
	}
	SystemUser = &User{
		Name:     "系统管理员",
		Username: viper.GetString("server.admin.user"),
		Password: viper.GetString("server.admin.pass"),
		Role:     user_iface.UserRole_USER_ROLE_SYSTEM,
		Status:   user_iface.UserStatus_USER_STATUS_NORMAL,
	}
	err := dao.Connection().Where(User{Username: SystemUser.Username, Role: SystemUser.Role}).FirstOrCreate(&SystemUser).Error
	if err != nil {
		panic(err)
	}
}
