package main

import (
	"context"
	"fmt"
	"supercrawler/common/pkg"
	"supercrawler/engine/config"
	"supercrawler/models"
	"testing"

	"github.com/spf13/viper"
)

func Test_PublishBook(t *testing.T) {
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	rdb.Publish(context.Background(), config.Queue.BookQueue, "https://www.ibiquge.info/168_168880/")
}

func Test_AllPublishBook(t *testing.T) {
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	ctx := context.Background()
	size := 1000
	for page := 0; ; page++ {
		var records []string
		models.NewBookClient(db).Connect().Select("src").Where("name=''").
			Offset(page * size).Limit(size).Find(&records)
		for _, record := range records {
			rdb.Publish(ctx, config.Queue.BookQueue, record)
		}
		if len(records) < config.Crawler.Parallelism {
			break
		}
		fmt.Println(page)
	}
}

func Test_PublishChapter(t *testing.T) {
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	rdb.Publish(context.Background(), config.Queue.ChapterQueue, "https://www.ibiquge.info/168_168880/65573011.html")
}

func Test_AllPublishChapter(t *testing.T) {
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	ctx := context.Background()
	size := 2000
	for page := 0; ; page++ {
		var records []string
		models.NewChapterClient(db).Connect().Select("src").Where("state!=?", models.BookStatusFinal).
			Offset(page * size).Limit(size).Find(&records)
		for _, record := range records {
			rdb.Publish(ctx, config.Queue.ChapterQueue, record)
		}
		if len(records) < config.Crawler.Parallelism {
			break
		}
		fmt.Println(page)
	}
}
