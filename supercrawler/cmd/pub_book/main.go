package main

import (
	"context"
	"flag"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"supercrawler/common/pkg"
	"supercrawler/engine"
	"supercrawler/engine/config"
	"supercrawler/models"
)

func main() {
	configPath := flag.String("c", "config.yaml", "config file path.")
	flag.Parse()
	if err := config.Read(*configPath); err != nil {
		panic(err)
	}
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	ctx := context.Background()
	for page := 0; ; page++ {
		list := make([]string, 0, config.Crawler.Parallelism)
		models.NewBookClient(db).Connect().Where("state=?", models.PENDING).
			Select("src").Limit(config.Crawler.Parallelism).Offset(page * config.Crawler.Parallelism).Find(&list)
		rdb.XAdd(ctx, &redis.XAddArgs{Stream: config.Crawler.Stream, Values: engine.EncodeParams(engine.BookLabel, list)})
		if len(list) < config.Crawler.Parallelism {
			break
		}
	}
}
