package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"supercrawler/common/pkg"
	"supercrawler/engine/config"
	"testing"
)

func Test_PublishBook(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	_ = viper.ReadInConfig()
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	rdb.Publish(context.Background(), config.BookQueue, "https://www.ibiquge.info/168_168880/")
}

func Test_PublishChapter(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	_ = viper.ReadInConfig()
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	rdb.Publish(context.Background(), config.ChapterQueue, "https://www.ibiquge.info/168_168880/65573011.html")
}

func Test_Config(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	_ = viper.ReadInConfig()
	cfg := config.Get()
	fmt.Printf("cfg: %+v", cfg)
}
