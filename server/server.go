package server

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"superserver/common/logger"
	"superserver/common/middleware"
	"superserver/docs"
	"superserver/internal/api"
	"superserver/internal/model/access"
	"superserver/internal/model/user"
	"superserver/pkg"
	"superserver/proto/user_iface"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func New() *Server {
	gin.SetMode(viper.GetString("mode"))
	switch viper.GetString("database.type") {
	case "mysql":

	}

	db, err := pkg.NewDbClient(viper.GetStringMap("database"), logger.NewGormLogger())
	if err != nil {
		panic(fmt.Sprintf("connect db failed: %s", err.Error()))
	}
	srv := &Server{engine: gin.New(), db: db}
	srv.middle = middleware.New()
	srv.engine.Use(
		srv.middle.CorsMiddleware(),
		srv.middle.RecoverMiddleware(),
		srv.middle.ContextMiddleware(),
	)
	return srv
}

type Server struct {
	engine *gin.Engine
	db     *gorm.DB
	app    *api.Api
	middle *middleware.Middleware
}

func (srv *Server) Route() {
	srv.app = api.New(srv.db, srv.engine)
	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Host = fmt.Sprintf(
			"%s:%d",
			viper.GetString("server.localhost"),
			viper.GetInt("server.port"),
		)
		srv.engine.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
	}
}

func (srv *Server) InitData() {
	user.NewDao(srv.db)

	// init system user
	user.SystemUser = &user.User{
		Name:     "系统管理员",
		Username: viper.GetString("admin.user"),
		Role:     user_iface.UserRole_USER_ROLE_SYSTEM,
		Status:   user_iface.UserStatus_USER_STATUS_NORMAL,
	}
	user.SystemUser.SetPassword(viper.GetString("admin.pass"))
	err := user.NewDao(srv.db).Upsert(user.SystemUser)
	if err != nil {
		panic(err)
	}

	// init access
	var accessList []access.Access
	for _, routeInfo := range srv.engine.Routes() {
		accessList = append(accessList, access.Access{
			Path: routeInfo.Path,
			// Method:  routeInfo.Method,
			Handler: routeInfo.Handler,
		})
	}
	err = access.NewDao(srv.db).BatchUpsert(accessList)
	if err != nil {
		panic(err)
	}
}

// Start 启动api服务
func (srv *Server) Start() error {
	// 路由注册
	srv.Route()
	// 初始化数据
	srv.InitData()
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
