package startup

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toolbox/common/middleware"
	"toolbox/common/system"
	"toolbox/internal/router"
	"toolbox/scheduler"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type WebapiService struct{}

func (srv *WebapiService) GetName() string {
	return "webapi"
}

func (srv *WebapiService) crontab() {
	crontab := scheduler.New()
	crontab.Start()
}

// Start 启动api服务
func (srv *WebapiService) Start(tags string) error {
	// start crontab schedule
	srv.crontab()

	// new engine
	engine := gin.New()

	// 测试路由
	router.TestRouteRegister(engine)

	// 中间件
	engine.Use(
		middleware.ContextMiddleware(),
		middleware.CorsMiddleware(),
		middleware.RecoverMiddleware(),
		middleware.AuthMiddleware(),
		middleware.LoggerMiddleware(),
	)

	// 路由注册
	router.BaseRouteRegister(engine)
	router.HandlerRouteRegister(engine)

	// 初始化服务上下文
	system.InitContextValue(tags)

	// http server
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", viper.GetString("server.port")),
		Handler:        engine,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return start(server)

}

// start 优雅启停
func start(server *http.Server) (err error) {
	errs := make(chan error, 1)
	go func() {
		if err = server.ListenAndServe(); err != nil {
			if err != nil {
				errs <- err
			}
		}
	}()
	scheduler := make(chan os.Signal, 1)
	signal.Notify(scheduler, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err = <-errs:
		break
	case s := <-scheduler:
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			err = server.Close()
		}
	}
	return err
}
