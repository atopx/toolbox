package engine

import (
	"context"
	"supercrawler/common/logger"
	"supercrawler/engine/book"
	"supercrawler/engine/chapter"
	"supercrawler/engine/config"
	"supercrawler/engine/sitemap"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Engine struct {
	db            *gorm.DB
	rdb           *redis.Client
	sitemapSpider *sitemap.Crawler
	bookSpider    *book.Crawler
	chapterSpider *chapter.Crawler
}

func New(db *gorm.DB, rdb *redis.Client) *Engine {
	return &Engine{
		db:            db,
		rdb:           rdb,
		sitemapSpider: sitemap.New(db),
		bookSpider:    book.New(db),
		chapterSpider: chapter.New(db),
	}
}

func (e *Engine) Start() {
	channels := []string{config.Queue.BookQueue, config.Queue.ChapterQueue, config.Queue.ChapterQueue}
	queue := e.rdb.Subscribe(context.Background(), channels...)
	logger.Info("start crawler engine success", zap.Strings("channels", channels))
	for event := range queue.Channel() {
		logger.Info("queue message", zap.String("queue", event.Channel), zap.String("src", event.Payload))
		switch event.Channel {
		case config.Queue.BookQueue:
			if err := e.bookSpider.Visit(event.Payload); err != nil {
				logger.Error("book visit error", zap.String("src", event.Payload), zap.Error(err))
			}
		case config.Queue.ChapterQueue:
			if err := e.chapterSpider.Visit(event.Payload); err != nil {
				logger.Error("chapter visit error", zap.String("src", event.Payload), zap.Error(err))
			}
		case config.Queue.SitemapQueue:
			if err := e.sitemapSpider.Visit(event.Payload); err != nil {
				logger.Error("sitemap visit error", zap.String("src", event.Payload), zap.Error(err))
			}
		}
	}
}
