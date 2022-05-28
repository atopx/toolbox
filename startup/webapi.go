package startup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"toolbox/common/middleware"
	"toolbox/internal/router"
)

type WebapiService struct{}

func (srv *WebapiService) GetName() string {
	return "webapi"
}

// Start 启动api服务
func (srv *WebapiService) Start() error {
	engine := gin.New()

	// 基础路由
	router.BaseRegister(engine)

	// 中间件
	engine.Use(
		middleware.RecoverMiddleware(),
		middleware.AuthMiddleware(),
		middleware.ContextMiddleware(),
		middleware.LoggerMiddleware(),
	)

	// 功能路由
	router.HandlerRegister(engine)

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
			break
		}
	}
	return err
}
