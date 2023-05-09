package scheduler

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"supercrawler/common/logger"
	"supercrawler/engine/book"
	"supercrawler/engine/config"
	"supercrawler/engine/sitemap"
	"supercrawler/models"
	"syscall"
)

type Scheduler struct {
	db         *gorm.DB
	rdb        *redis.Client
	queue      *redis.PubSub
	cron       *cron.Cron
	siteSpider *sitemap.Crawler
	bookSpider *book.Crawler
}

func New(db *gorm.DB, rdb *redis.Client) *Scheduler {
	return &Scheduler{
		db:         db,
		rdb:        rdb,
		queue:      rdb.Subscribe(context.Background(), config.QueueName),
		cron:       cron.New(),
		siteSpider: sitemap.New(db),
		bookSpider: book.New(db),
	}
}

func (s *Scheduler) register() {
	_, _ = s.cron.AddFunc(config.SitemapCron, func() {
		s.siteSpider.Start()
	})
	_, _ = s.cron.AddFunc(config.BookCron, func() {
		// 查询NEW, 发送任务
		bookClient := models.NewBookClient(s.db)
		var count int64
		bookClient.Connect().Where("state=?", models.BookStatusRunning).Count(&count)
		if count < 10 {
			var srcs []string
			bookClient.Connect().Where("state=?", models.BookStatusNew).
				Select("src").Order("id").Limit(10).Find(&srcs)
			bookClient.Connect().Where("src in (?)", srcs).UpdateColumn("state", models.BookStatusRunning)
			for _, src := range srcs {
				s.rdb.Publish(context.Background(), config.QueueName, src)
			}
		}
	})
	s.cron.Start()
	logger.Info("start cron success")
}

func (s *Scheduler) Start() {
	s.register()
	go func() {
		for msg := range s.queue.Channel(redis.WithChannelSize(1)) {
			logger.Info("queue message", zap.String("queue", msg.Channel), zap.String("payload", msg.Payload))
			_ = s.bookSpider.Visit(msg.Payload)
			s.bookSpider.Wait()
		}
	}()
	manager := make(chan os.Signal, 1)
	signal.Notify(manager, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case op := <-manager:
		switch op {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.cron.Stop()
			_ = s.queue.Close()
		}
	}
}
