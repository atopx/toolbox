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
	"superserver/internal/model/role"
	"superserver/internal/model/user"
	"superserver/internal/scheduler"
	"superserver/pkg"
	"superserver/proto/common_iface"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	files "github.com/swaggo/files"
	swagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func New() *Server {
	gin.SetMode(viper.GetString("server.mode"))
	db, err := pkg.NewDbClient(viper.GetStringMap("database"), logger.NewGormLogger())
	if err != nil {
		panic(fmt.Sprintf("connect db failed: %s", err.Error()))
	}
	srv := &Server{engine: gin.New(), db: db, middle: middleware.New()}
	srv.crontab = scheduler.New(db)
	srv.listen = fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
	srv.engine.Use(srv.middle.CorsMiddleware(), srv.middle.RecoverMiddleware(), srv.middle.ContextMiddleware())
	return srv
}

type Server struct {
	engine  *gin.Engine
	db      *gorm.DB
	app     *api.Api
	middle  *middleware.Middleware
	listen  string
	crontab *cron.Cron
}

func (srv *Server) Route() {
	srv.app = api.New(srv.db, srv.engine)
	if gin.Mode() == gin.DebugMode {
		docs.SwaggerInfo.Host = srv.listen
		srv.engine.GET("/swagger/*any", swagger.WrapHandler(files.Handler))
		logger.System("swagger web: http://%s/swagger/index.html", srv.listen)
	}
}

func (srv *Server) InitData() {
	// init roles
	tx := srv.db.Begin()
	err := role.NewDao(tx).LoadSystemRole()
	if err != nil {
		tx.Rollback()
		panic(fmt.Errorf("init system role failed: %s", err.Error()))
	}

	err = role.NewDao(tx).LoadDefaultRole()
	if err != nil {
		tx.Rollback()
		panic(fmt.Errorf("init default role failed: %s", err.Error()))
	}
	fmt.Println(role.SystemRole)
	// init user
	err = user.NewDao(srv.db).LoadSystemUser(role.SystemRole.Id)
	if err != nil {
		tx.Rollback()
		panic(fmt.Errorf("init system user failed: %s", err.Error()))
	}

	// init access
	var accessList []access.Access
	for _, routeInfo := range srv.engine.Routes() {
		accessList = append(accessList, access.Access{
			Path:    routeInfo.Path,
			Method:  common_iface.AccessMethod(common_iface.AccessMethod_value[routeInfo.Method]),
			Handler: routeInfo.Handler,
		})
	}
	err = access.NewDao(srv.db).BatchUpsert(accessList)
	if err != nil {
		tx.Rollback()
		panic(fmt.Errorf("init access failed: %s", err.Error()))
	}

	tx.Commit()
}

// Start 启动api服务
func (srv *Server) Start() error {
	// 启动后台任务
	srv.crontab.Start()
	// 路由注册
	srv.Route()
	// 初始化数据
	srv.InitData()
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
	scheduler := make(chan os.Signal, 1)
	signal.Notify(scheduler, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case err = <-errs:
		break
	case s := <-scheduler:
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			srv.crontab.Stop()
			err = server.Close()
		}
	}
	return err
}
