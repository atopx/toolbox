package main

import (
	"context"
	"fmt"
	"supercrawler/common/pkg"
	"supercrawler/engine/config"
	"supercrawler/models"
	"testing"

	"github.com/spf13/viper"
	"gorm.io/gorm"
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
	var records []models.Book
	models.NewBookClient(db).Connect().Order("id").Select("src").
		Where("name='' and state!=?", models.BookStatusFinal).
		FindInBatches(&records, 1000, func(tx *gorm.DB, batch int) error {
			for _, record := range records {
				rdb.Publish(ctx, config.Queue.BookQueue, record.Src)
			}
			fmt.Println()
			return nil
		})
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
	var records []models.Chapter
	ctx := context.Background()
	models.NewChapterClient(db).Connect().Order("id").Select("src").
		Where("state!=?", models.BookStatusFinal).
		FindInBatches(&records, 1000, func(tx *gorm.DB, batch int) error {
			for _, record := range records {
				rdb.Publish(ctx, config.Queue.ChapterQueue, record.Src)
			}
			return nil
		})
}
