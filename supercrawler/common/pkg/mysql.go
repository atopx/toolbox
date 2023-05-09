package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

func NewDbClient(cfg map[string]interface{}, logger glog.Interface) (*gorm.DB, error) {
	if logger == nil {
		logger = defaultDbLogger()
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg["user"], cfg["pass"], cfg["host"], cfg["port"], cfg["name"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                func() time.Time { return time.Now().Local() },
		Logger:                 logger,
		PrepareStmt:            true,
		CreateBatchSize:        800,
	})
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(20)
	sqlDb.SetConnMaxLifetime(time.Hour)
	return db, nil
}

func defaultDbLogger() glog.Interface {
	return glog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		glog.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  glog.Silent, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
}
