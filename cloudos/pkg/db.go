package pkg

import (
	"cloudos/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewMysqlClient(cfg *config.MysqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name,
	)
	ormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NowFunc:                func() time.Time { return time.Now().Local() },
		PrepareStmt:            true,
		CreateBatchSize:        1000,
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
	}
	if cfg.Debug {
		ormConfig.Logger = debugLogger()
	}
	db, err := gorm.Open(mysql.Open(dsn), ormConfig)
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func debugLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
}
