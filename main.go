package main

import (
	"context"
	"flag"
	"log"

	"superserver/common/logger"
	"superserver/server"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	_ "superserver/docs"
)

// @title           superserver
// @version         0.1.0
// @description     集成网盘、多媒体、Webdav、服务器监控等一体的超级服务器后台

// @contact.name   atopx
// @contact.email  3940422@qq.com
// @contact.url    https://github.com/atopx/superserver.git

func main() {
	configPath := flag.String("c", "configs/develop.yaml", "config file path.")
	flag.Parse()
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("server.mode", "release")
		viper.Set("server.listen_addr", "127.0.0.1")
		viper.Set("server.listen_port", 8000)
		viper.Set("server.admin_user", "superuser")
		viper.Set("server.admin_pass", "superuser")
		log.Println("user default config:", err)
	}

	// 日志初始化
	var loglevel zapcore.Level
	if err := logger.Setup(zapcore.InfoLevel.String()); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}
	log.Printf("start server by [%s].\n", loglevel.String())

	// 启动服务
	if err := server.New().Start(); err != nil {
		logger.Panic(context.TODO(), "start server panic", zap.Error(err))
	}
}
