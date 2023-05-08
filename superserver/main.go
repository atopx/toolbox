package main

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"superserver/common/logger"
	"superserver/internal/controller/novel"
	"superserver/server"
)

// @title           superserver
// @version         0.1.0
// @host			localhost:18000
// @description     集成网盘、多媒体、Webdav、服务器监控等一体的超级服务器后台

// @contact.name   atopx
// @contact.email  3940422@qq.com
// @contact.url    https://github.com/atopx/superserver.git

func main() {
	// 配置初始化
	config := flag.String("c", "config.yaml", "config file path.")
	worker := flag.String("w", "", "start a worker")
	flag.Parse()
	viper.SetConfigFile(*config)
	if err := viper.ReadInConfig(); err != nil {
		viper.Set("server.mode", "release")
		viper.Set("server.host", "127.0.0.1")
		viper.Set("server.port", 18000)
		viper.Set("server.loglevel", "info")
		viper.Set("admin.user", "superuser")
		viper.Set("admin.pass", "superuser")
		log.Println("using default server config:")
		for key, value := range viper.GetStringMap("server") {
			log.Printf(" - %s: %v\n", key, value)
		}
	}

	// 日志初始化
	if err := logger.Setup(viper.GetString("server.loglevel")); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}

	// 启动服务
	switch *worker {
	case "novel":
		novel.New().Listen()
	default:
		if err := server.New().Start(); err != nil {
			log.Panicf("start server failed: %s", err.Error())
		}
	}
}
