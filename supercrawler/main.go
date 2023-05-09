package main

import (
	"flag"
	"log"
	"supercrawler/common/logger"
	"supercrawler/common/pkg"
	"supercrawler/engine/scheduler"

	"github.com/spf13/viper"
)

func main() {
	configPath := flag.String("c", "config.yaml", "config file path.")
	flag.Parse()
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 日志初始化
	if err := logger.Setup(viper.GetString("server.loglevel")); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}
	db, err := pkg.NewDbClient(viper.GetStringMap("mysql"), nil)
	if err != nil {
		panic(err)
	}
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	scheduler.New(db, rdb).Start()
}
