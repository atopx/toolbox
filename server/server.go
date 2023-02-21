package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"superserver/common/logger"
	"superserver/common/middleware"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Server struct {
	engine *gin.Engine
}

func New() *Server {
	gin.SetMode(viper.GetString("mode"))
	srv := &Server{engine: gin.New()}
	srv.engine.Use(
		middleware.CorsMiddleware(),
		middleware.RecoverMiddleware(),
		middleware.LoggerMiddleware(),
	)
	return srv
}

// Start 启动api服务
func (srv *Server) Start() error {
	// 路由注册
	srv.Route()

	// http server
	listen := fmt.Sprintf("%s:%d", viper.GetString("server.addr"), viper.GetInt("server.port"))
	server := &http.Server{
		Addr:           listen,
		Handler:        srv.engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.System("server listen: http://%s", listen)
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
