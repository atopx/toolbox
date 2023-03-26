package server

import (
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
	srv.listen = fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
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
	listen string
}

func (srv *Server) Route() {

}

//func (srv *Server) InitData() {
//	// init roles
//	tx := srv.db.Begin()
//	err := role.NewDao(tx).LoadSystemRole()
//	if err != nil {
//		tx.Rollback()
//		panic(fmt.Errorf("init system role failed: %s", err.Error()))
//	}
//
//	err = role.NewDao(tx).LoadDefaultRole()
//	if err != nil {
//		tx.Rollback()
//		panic(fmt.Errorf("init default role failed: %s", err.Error()))
//	}
//
//	// init admin user
//	err = user.NewDao(srv.db).LoadSystemUser()
//	if err != nil {
//		tx.Rollback()
//		panic(fmt.Errorf("init system user failed: %s", err.Error()))
//	}
//
//	// bind user and role
//	userRole := user_role.UserRoleRef{RoleId: role.SystemRole.Id, UserId: user.SystemUser.Id}
//	user_role.NewDao(srv.db).Connection().FirstOrCreate(&userRole)
//
//	// init access
//	var accessList []access.Access
//	for _, routeInfo := range srv.engine.Routes() {
//		record := access.Access{
//			Path:    routeInfo.Path,
//			Method:  routeInfo.Method,
//			Handler: routeInfo.Handler,
//		}
//		accessList = append(accessList, record)
//	}
//	if err = access.NewDao(srv.db).BatchUpsert(accessList); err != nil {
//		tx.Rollback()
//		panic(fmt.Errorf("init access failed: %s", err.Error()))
//	}
//	tx.Commit()
//}

// Start 启动api服务
func (srv *Server) Start() error {
	// 路由注册
	api.Register(srv.engine)
	// 初始化数据
	//srv.InitData()
	// http server
	server := &http.Server{
		Addr:           srv.listen,
		Handler:        srv.engine,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.System("server listen: http://%s", srv.listen)
	return srv.start(server)
}

// start 优雅启停
func (srv *Server) start(server *http.Server) (err error) {
	errs := make(chan error, 1)
	go func() {
		if err = server.ListenAndServe(); err != nil {
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
			err = server.Close()
		}
	}
	return err
}
