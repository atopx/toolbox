package pkg

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewRedisClient(cfg map[string]interface{}) *redis.Client {
	var password string
	if value := cfg["password"]; value != nil {
		password = value.(string)
	}

	return redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", cfg["host"], cfg["port"]),
		Password:        password,
		DB:              cfg["db"].(int),
		ConnMaxLifetime: time.Hour,
	})
}
