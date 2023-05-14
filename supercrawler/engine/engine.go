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
	e.rdb.XGroupCreateMkStream(ctx, config.Crawler.Stream, config.Crawler.Group, "0")
	e.rdb.XGroupCreateConsumer(ctx, config.Crawler.Stream, config.Crawler.Group, config.Crawler.Stream)
	args := &redis.XReadGroupArgs{
		Group:    config.Crawler.Group,
		Consumer: config.Crawler.Group,
		Streams:  []string{config.Crawler.Stream, ">"},
		Count:    24,
		Block:    0,
		NoAck:    true,
	}
	for {
		reader := e.rdb.XReadGroup(ctx, args)
		queues, err := reader.Result()
		if err != nil {
			logger.Panic("read group failed", zap.Error(err))
		}
		message := queues[0].Messages[0]
		label, data := DecodeParams(message.Values)
		logger.Info("consumer message", zap.String("id", message.ID), zap.String("label", label))
		var spider Spider
		switch label {
		case BookLabel:
			spider = e.bookSpider
		case ChapterLabel:
			spider = e.chapterSpider
		case SitemapLabel:
			spider = e.sitemapSpider
		}
		if spider != nil {
			for _, src := range data {
				_ = spider.Visit(src)
			}
			spider.Wait()
		}
		logger.Info("consumer success", zap.String("id", message.ID), zap.String("id", message.ID), zap.String("label", label))
		e.rdb.XAck(ctx, config.Crawler.Stream, config.Crawler.Group, message.ID)
		e.rdb.XDel(ctx, config.Crawler.Stream, message.ID)
	}
}
