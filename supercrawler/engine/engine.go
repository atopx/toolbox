package engine

import (
	"context"
	"go.uber.org/zap"
	"supercrawler/common/logger"
	"supercrawler/engine/book"
	"supercrawler/engine/chapter"
	"supercrawler/engine/config"
	"supercrawler/engine/sitemap"

	"github.com/redis/go-redis/v9"
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

const (
	BookLabel    = "book"
	ChapterLabel = "chapter"
	SitemapLabel = "sitemap"
)

type Spider interface {
	Visit(url string) error
	Wait()
}

func (e *Engine) Start() {
	ctx := context.Background()
	args := &redis.XReadArgs{Streams: []string{config.Crawler.Queue, "$"}, Count: 1, Block: 0}
	for {
		reader := e.rdb.XRead(ctx, args)
		queues, err := reader.Result()
		if err != nil {
			logger.Panic("read group failed", zap.Error(err))
		}
		message := queues[0].Messages[0]
		logger.Info("consumer message", zap.Any("data", message))
		var spider Spider
		switch message.Values["label"].(string) {
		case BookLabel:
			spider = e.bookSpider
		case ChapterLabel:
			spider = e.chapterSpider
		case SitemapLabel:
			spider = e.sitemapSpider
		}
		if spider != nil {
			if err = spider.Visit(message.Values["src"].(string)); err != nil {
				logger.Error("visit error", zap.Error(err))
			} else {
				spider.Wait()
				logger.Info("consumer success")
			}
		}
		e.rdb.XDel(ctx, config.Crawler.Queue, message.ID)
	}
}
