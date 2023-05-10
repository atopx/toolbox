package main

import (
	"flag"
	"log"
	"supercrawler/common/logger"
	"supercrawler/common/pkg"
	"supercrawler/engine"
	"supercrawler/engine/config"

	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("c", "config.yaml", "config file path.")
	flag.Parse()
	if err := config.Read(*configPath); err != nil {
		panic(err)
	}

	// 日志初始化
	if err := logger.Setup(viper.GetString("server.loglevel")); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}

	// init db
	db := pkg.NewDbClient(viper.GetStringMap("mysql"))
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))

	// start engine
	engine.New(db, rdb).Start()
}
