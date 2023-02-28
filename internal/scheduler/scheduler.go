package scheduler

import (
	"superserver/internal/scheduler/novels"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *cron.Cron {
	var crontab = cron.New()

	// novel doenloader
	novelTask := novels.New(db)
	crontab.AddFunc(novelTask.Cron(), novelTask.Start)

	return crontab
}
