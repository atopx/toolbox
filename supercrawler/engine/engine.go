package engine

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"supercrawler/common/logger"
	"supercrawler/engine/book"
	"supercrawler/engine/chapter"
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

type Message struct {
	Label string `json:"label"`
	Src   string `json:"src"`
}

func (e *Engine) Start() {
	ctx := context.Background()
	stream := "super-crawler"
	group := "crawler-group"
	consumer := "crawler-consumer"
	if err := e.rdb.XGroupCreateMkStream(ctx, stream, group, "0").Err(); err != nil {
		logger.Panic("create group failed", zap.Error(err))
	}
	for {
		cmd := e.rdb.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    group,
			Consumer: consumer,
			Streams:  []string{stream, ">"},
			Count:    1,
		})
		queues, err := cmd.Result()
		if err != nil {
			fmt.Println(err)
			logger.Panic("read group failed", zap.Error(err))
		}
		message := queues[0].Messages[0]
		label := message.Values["label"].(string)
		src := message.Values["src"].(string)
		logger.Info("queue message", zap.String("key", message.ID), zap.String("label", label), zap.String("src", src))
		switch message.Values["label"].(string) {
		case "book":
			if err = e.bookSpider.Visit(src); err != nil {
				logger.Error("book visit error", zap.String("src", src), zap.Error(err))
				continue
			}
			err = e.bookSpider.Visit(src)
			e.bookSpider.Wait()
		case "chapter":
			if err = e.chapterSpider.Visit(src); err != nil {
				logger.Error("chapter visit error", zap.String("src", src), zap.Error(err))
				continue
			}
			e.chapterSpider.Wait()
		case "sitemap":
			if err = e.sitemapSpider.Visit(src); err != nil {
				logger.Error("sitemap visit error", zap.String("src", src), zap.Error(err))
				continue
			}
			e.sitemapSpider.Wait()
		}
		e.rdb.XAck(ctx, stream, group, message.ID)
	}
}
