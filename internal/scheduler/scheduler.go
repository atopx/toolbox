package scheduler

import (
	"superserver/internal/scheduler/novels"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *cron.Cron {
	var crontab = cron.New()

	// novel downloader task
	novelTask := novels.New(db)
	_, _ = crontab.AddFunc(novelTask.Cron(), novelTask.Start)

	return crontab
}
