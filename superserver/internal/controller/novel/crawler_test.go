package novel

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"superserver/pkg"
	"testing"
)

func TestCrawler_Pub(t *testing.T) {
	viper.SetConfigFile("../../../config.yaml")
	viper.ReadInConfig()

	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	addr := "https://www.ibiquge.info/169_169240/"
	rdb.Publish(context.Background(), "crawler-novel", addr)
}

func TestCrawler_Sub(t *testing.T) {
	viper.SetConfigFile("../../../config.yaml")
	viper.ReadInConfig()

	rdb := pkg.NewRedisClient(viper.GetStringMap("redis"))
	addr := "https://www.ibiquge.info/169_169240/"
	key := fmt.Sprintf("novel::%s", addr)
	result, _ := rdb.HGetAll(context.Background(), key).Result()
	fmt.Println(result["name"], result["author"])
	rdb.Del(context.Background(), key)
}
