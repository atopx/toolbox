package scheduler

import (
	"toolbox/scheduler/gfwrules"

	"github.com/robfig/cron/v3"
)

func New() *cron.Cron {
	var crontab = cron.New()

	// gfw downloader
	crontab.AddFunc(gfwrules.Cron, gfwrules.GFWDownloader)

	return crontab
}
