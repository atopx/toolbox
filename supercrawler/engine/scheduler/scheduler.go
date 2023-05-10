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
	"supercrawler/common/utils"
	"supercrawler/engine/book"
	"supercrawler/engine/chapter"
	"supercrawler/engine/config"
	"supercrawler/engine/sitemap"
	"supercrawler/models"
	"syscall"
)

type Scheduler struct {
	db            *gorm.DB
	rdb           *redis.Client
	subscribe     *redis.PubSub
	queue         <-chan *redis.Message
	cron          *cron.Cron
	siteSpider    *sitemap.Crawler
	bookSpider    *book.Crawler
	chapterSpider *chapter.Crawler
}

func New(db *gorm.DB, rdb *redis.Client) *Scheduler {
	return &Scheduler{
		db:            db,
		rdb:           rdb,
		subscribe:     rdb.Subscribe(context.Background(), config.BookQueue, config.ChapterQueue),
		cron:          cron.New(),
		siteSpider:    sitemap.New(db),
		bookSpider:    book.New(db),
		chapterSpider: chapter.New(db),
	}
}

func (s *Scheduler) register() {
	_, _ = s.cron.AddFunc(config.SitemapCron, func() { s.siteSpider.Start() })

	_, _ = s.cron.AddFunc(config.CrawlerCron, func() {
		if len(s.queue) < config.Parallelism {
			bookClient := models.NewBookClient(s.db)
			var count int64
			bookClient.Connect().Where("state=?", models.BookStatusNew).Count(&count)
			if count > 0 {
				// 发送book任务
				var srcs []string
				bookClient.Connect().Where("state=?", models.BookStatusNew).Select("src").Order("id").Limit(config.Threshold).Find(&srcs)
				bookClient.Connect().Where("src in (?)", srcs).UpdateColumn("state", models.BookStatusRunning)
				s.rdb.Publish(context.Background(), config.BookQueue, s.encodeMessage(srcs))
			} else {
				// 发送chapter任务
				var srcs []string
				chapterClient := models.NewChapterClient(s.db)
				chapterClient.Connect().Where("state=?", models.BookStatusNew).Select("src").Order("id").Limit(config.Threshold).Find(&srcs)
				chapterClient.Connect().Where("src in (?)", srcs).UpdateColumn("state", models.BookStatusRunning)
				s.rdb.Publish(context.Background(), config.ChapterQueue, s.encodeMessage(srcs))
			}
		}
	})
	s.cron.Start()
	logger.Info("start cron success")
}

func (s *Scheduler) encodeMessage(srcs []string) string {
	return utils.StringMarshal(srcs)
}

func (s *Scheduler) decodeMessage(message string) (srcs []string) {
	utils.StringUnmarshal(message, &srcs)
	return srcs
}

func (s *Scheduler) Start() {
	s.queue = s.subscribe.Channel()
	s.register()
	go func() {
		for event := range s.queue {
			srcs := s.decodeMessage(event.Payload)
			logger.Info("queue message", zap.String("queue", event.Channel), zap.Strings("srcs", srcs))
			switch event.Channel {
			case config.BookQueue:
				for _, src := range srcs {
					_ = s.bookSpider.Visit(src)
				}
				s.bookSpider.Wait()
			case config.ChapterQueue:
				for _, src := range srcs {
					_ = s.chapterSpider.Visit(src)
				}
				s.chapterSpider.Wait()
			}
		}
	}()
	manager := make(chan os.Signal, 1)
	signal.Notify(manager, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case op := <-manager:
		switch op {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.cron.Stop()
			_ = s.subscribe.Close()
		}
	}
}
