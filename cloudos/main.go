package main

import (
	"cloudos/common/logger"
	"cloudos/common/system"
	"cloudos/internal/worker/workmail"
	"cloudos/server"
	"flag"
	"log"
)

// @title           cloudos
// @version         0.1.0
// @host			localhost:18000
// @description     集成网盘、多媒体、Webdav、服务器监控等一体的超级服务器后台

// @contact.name   atopx
// @contact.email  3940422@qq.com
// @contact.url    https://github.com/atopx/cloudos.git

func main() {
	// 配置初始化
	path := flag.String("c", "config.yaml", "config file path.")
	wkmail := flag.String("wkmail", "", "send work mail, params: noteId/datetime, example: -wkmail 123/2006-01-02 15:04:05")
	flag.Parse()

	system.SetHandler(*path)

	// 日志初始化
	if err := logger.Setup(system.GetHandler().Config.Server.Loglevel); err != nil {
		log.Panicf("logger setup failed: %s", err.Error())
	}

	if *wkmail != "" {
		if err := workmail.Scheduler(*wkmail); err != nil {
			panic(err)
		}
		return
	}

	// 启动服务
	if err := server.New().Start(); err != nil {
		log.Panicf("start server failed: %s", err.Error())
	}
}
