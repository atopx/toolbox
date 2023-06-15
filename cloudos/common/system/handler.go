package system

import (
	"cloudos/config"
	"cloudos/pkg"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Handler struct {
	Config config.Config
	Db     *gorm.DB
	Rdb    *redis.Client
}

var handler *Handler
var once sync.Once

func SetHandler(path string) {
	once.Do(func() {
		viper.SetConfigFile(path)
		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		handler = new(Handler)
		if err = viper.Unmarshal(&handler.Config); err != nil {
			panic(err)
		}

		handler.Db, err = pkg.NewMysqlClient(handler.Config.Mysql)
		if err != nil {
			panic(err)
		}

		handler.Rdb = pkg.NewRedisClient(handler.Config.Redis)
	})
}

func GetHandler() *Handler {
	return handler
}
