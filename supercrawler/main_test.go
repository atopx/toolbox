package main_test

import (
	"bytes"
	"context"
	"os"
	"supercrawler/common/pkg"
	"supercrawler/engine"
	"supercrawler/engine/config"
	"supercrawler/models"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/spf13/viper"
)

func publish(value interface{}) error {
	if err := config.Read("config.yaml"); err != nil {
		return err
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	return rdb.XAdd(context.Background(), &redis.XAddArgs{Stream: config.Crawler.Stream, Values: value}).Err()
}

func Test_Publish_sitemap(t *testing.T) {
	if err := publish(map[string]interface{}{
		"label": engine.SitemapLabel,
		"src":   "https://www.ibiquge.info/api/sitemap.xml",
	}); err != nil {
		t.Fatal(err)
	}
}

func Test_Publish_book(t *testing.T) {
	bookSrc := "https://www.ibiquge.info/131_131180/"
	if err := publish(engine.EncodeParams(engine.BookLabel, []string{bookSrc})); err != nil {
		t.Fatal(err)
	}
}

func Test_Publish_chapter(t *testing.T) {
	bookSrc := "https://www.ibiquge.info/131_131180/%"
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	var data []string
	models.NewChapterClient(db).Connect().Select("src").Where("src like ?", bookSrc).Find(&data)
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	params := engine.EncodeParams(engine.ChapterLabel, data)
	rdb.XAdd(context.Background(), &redis.XAddArgs{Stream: config.Crawler.Stream, Values: params})
}

func Test_Export(t *testing.T) {
	bookSrc := "https://www.ibiquge.info/131_131180/%"
	if err := config.Read("config.yaml"); err != nil {
		t.Fatal(err)
	}
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	var chapters []models.Chapter
	models.NewChapterClient(db).Connect().Select("title", "content").
		Where("src like ?", bookSrc).Order("code").Find(&chapters)
	var buf bytes.Buffer
	buf.WriteString("《万象之王》\n")
	buf.WriteString("作者: 天蚕土豆\n")
	for _, chapter := range chapters {
		buf.WriteString(chapter.Title)
		buf.WriteString("\n")
		buf.WriteString(chapter.Content)
		buf.WriteString("\n")
	}
	_ = os.WriteFile(time.Now().Local().Format("万象之王20060102150405.txt"), buf.Bytes(), 0700)
}
