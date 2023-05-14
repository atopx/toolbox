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

func Test_Publish(t *testing.T) {
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	err := rdb.XAdd(context.Background(), &redis.XAddArgs{
		Stream: config.Crawler.Queue,
		Values: map[string]interface{}{
			"label": engine.SitemapLabel,
			"src":   "https://www.ibiquge.info/api/sitemap.xml",
		},
	}).Err()
	if err != nil {
		t.Fatal(err)
	}
}
