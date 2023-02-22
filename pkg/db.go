package pkg

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glog "gorm.io/gorm/logger"
)

func NewDbClient(dbpath string, logger glog.Interface) (*gorm.DB, error) {
	if dbpath == "" {
		dbpath = "super.data"
	}

	if logger == nil {
		logger = glog.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
			glog.Config{
				SlowThreshold:             time.Second, // 慢 SQL 阈值
				LogLevel:                  glog.Silent, // 日志级别
				IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,        // 禁用彩色打印
			},
		)
	}
	return gorm.Open(sqlite.Open(dbpath), &gorm.Config{
		SkipDefaultTransaction: false,
		NowFunc:                func() time.Time { return time.Now().Local() },
		Logger:                 logger,
		PrepareStmt:            true,
		CreateBatchSize:        1000,
	})
}
