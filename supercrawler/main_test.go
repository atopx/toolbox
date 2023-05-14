package main_test

import (
	"context"
	"supercrawler/common/pkg"
	"supercrawler/engine"
	"supercrawler/engine/config"
	"testing"

	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
)

func publish(value map[string]interface{}) error {
	if err := config.Read("config.yaml"); err != nil {
		return err
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	return rdb.XAdd(context.Background(), &redis.XAddArgs{Stream: config.Crawler.Queue, Values: value}).Err()
}

func Test_Publish_sitemap(t *testing.T) {
	if err := publish(map[string]interface{}{
		"label": engine.BookLabel,
		"src":   "https://www.ibiquge.info/api/sitemap.xml",
	}); err != nil {
		t.Fatal(err)
	}
}

func Test_Publish_book(t *testing.T) {
	if err := publish(map[string]interface{}{
		"label": engine.BookLabel,
		"src":   "https://www.ibiquge.info/157_157430/",
	}); err != nil {
		t.Fatal(err)
	}
}

func Test_Publish_chapter(t *testing.T) {
	if err := publish(map[string]interface{}{
		"label": engine.BookLabel,
		"src":   "https://www.ibiquge.info/157_157430/60934730.html",
	}); err != nil {
		t.Fatal(err)
	}
}
