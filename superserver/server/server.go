package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"superserver/common/logger"
	"superserver/common/middleware"
	"superserver/internal/api"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func New() *Server {
	gin.SetMode(viper.GetString("server.mode"))
	srv := &Server{engine: gin.New(), middle: middleware.New()}
	srv.engine.Use(
		srv.middle.CorsMiddleware(),
		srv.middle.RecoverMiddleware(),
		srv.middle.ContextMiddleware(),
		srv.middle.AuthMiddleware(),
	)
	return srv
}

type Server struct {
	engine *gin.Engine
	middle *middleware.Middleware
	server *http.Server
}

func (srv *Server) InitData() {
	ctx := context.Background()
	initAccess(ctx, srv.engine.Routes())
	initSystemUser(ctx)
	initInternalRoles(ctx)
}

// Start 启动api服务
func (srv *Server) Start() error {
	api.Register(srv.engine)
	srv.InitData()
	srv.server = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port")),
		Handler:        srv.engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.System("server listen: http://%s", srv.server.Addr)
	return srv.loop()
}

// start 优雅启停
func (srv *Server) loop() (err error) {
	errs := make(chan error, 1)
	go func() {
		if err = srv.server.ListenAndServe(); err != nil {
			if err != nil {
				errs <- err
			}
		}
	}()
	manager := make(chan os.Signal, 1)
	signal.Notify(manager, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err = <-errs:
		break
	case s := <-manager:
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			err = srv.server.Close()
		}
	}
	return err
}
