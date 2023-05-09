package main

import (
	"context"
	"github.com/spf13/viper"
	"supercrawler/common/pkg"
	"supercrawler/engine/config"
	"testing"
)

func Test_Publish(t *testing.T) {
	viper.SetConfigFile("config.yaml")
	_ = viper.ReadInConfig()
	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	rdb.Publish(context.Background(), config.QueueName, "https://www.ibiquge.info/168_168880/65573011.html")
}
