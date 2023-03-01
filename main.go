package main

import (
	"flag"
	"log"

	"superserver/common/logger"
	"superserver/server"

	"github.com/spf13/viper"

	_ "superserver/docs"
)

// @title           superserver
// @version         0.1.0
// @description     集成网盘、多媒体、Webdav、服务器监控等一体的超级服务器后台

// @contact.name   atopx
// @contact.email  3940422@qq.com
// @contact.url    https://github.com/atopx/superserver.git

func main() {
	// 配置初始化
	configPath := flag.String("c", "config.yaml", "config file path.")
	flag.Parse()
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("server.mode", "release")
		viper.Set("server.host", "127.0.0.1")
		viper.Set("server.port", 8000)
		viper.Set("server.loglevel", "info")
		viper.Set("admin.user", "superuser")
		viper.Set("admin.pass", "superuser")
		log.Println("useing default server config:")
		for key, value := range viper.GetStringMap("server") {
			log.Printf(" - %s: %v\n", key, value)
		}
	}

	// 日志初始化
	if err := logger.Setup(viper.GetString("server.loglevel")); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}

	// 启动服务
	if err := server.New().Start(); err != nil {
		log.Panicf("start server failed: %s", err.Error())
	}
}
