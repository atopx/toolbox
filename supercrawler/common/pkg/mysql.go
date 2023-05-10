package pkg

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDbClient(cfg map[string]interface{}) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg["user"], cfg["pass"], cfg["host"], cfg["port"], cfg["name"],
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		NowFunc:                func() time.Time { return time.Now().Local() },
		Logger:                 defaultDbLogger(),
		PrepareStmt:            true,
		CreateBatchSize:        500,
	})
	if err != nil {
		panic(err)
	}
	handle, err := db.DB()
	if err != nil {
		panic(err)
	}
	handle.SetMaxIdleConns(10)
	handle.SetMaxOpenConns(50)
	handle.SetConnMaxLifetime(time.Hour)
	return db
}

func defaultDbLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	)
}
