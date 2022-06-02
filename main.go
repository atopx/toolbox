package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"
	"toolbox/common/logger"
	"toolbox/common/system"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "toolbox/docs"
	_ "toolbox/startup"
)

// @title           toolbox-api
// @version         1.0
// @description     工具箱api

// @contact.name   atopx
// @contact.url    https://github.com/atopx/toolbox.git
// @contact.email  3940422@qq.com

var tags string

func main() {
	// 系统初始化
	services := system.AllService()
	if len(services) == 0 {
		log.Panicf("no register service.")
	}
	configPath := *flag.String("c", "configs/develop.yaml", "config file path.")
	serviceName := *flag.String("s", services[0], fmt.Sprintf("service start-up: [ %s ]", strings.Join(services, " | ")))
	flag.Parse()

	service := system.GetService(serviceName)
	if service == nil {
		log.Panicf("no register service.")
	}

	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("load config error: %s", err)
	}
	// 日志初始化
	var loglevel zapcore.Level
	mode := viper.GetString("server.env")
	switch mode {
	case "dev", "develop", "test":
		gin.SetMode(gin.DebugMode)
		loglevel = zap.DebugLevel
	case "prod", "product", "release":
		gin.SetMode(gin.ReleaseMode)
		loglevel = zap.InfoLevel
	default:
		gin.SetMode(gin.ReleaseMode)
		loglevel = zap.InfoLevel
	}
	if err := logger.Setup(loglevel.String()); err != nil {
		log.Panicf("invalid run mode: %s", mode)
	}
	log.Printf("start server by %s[%s].\n", mode, loglevel.String())

	// 启动服务
	if err := service.Start(tags); err != nil {
		logger.Panic(context.TODO(), "start server panic", zap.Error(err))
	}
}
