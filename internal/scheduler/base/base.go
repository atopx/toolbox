package base

import (
	"context"

	"gorm.io/gorm"
)

type Scheduler struct {
	Db      *gorm.DB
	Context context.Context
}

func NewScheduler(db *gorm.DB) Scheduler {
	return Scheduler{Db: db}
}
