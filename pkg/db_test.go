package pkg

import (
	"testing"

	"github.com/spf13/viper"
)

func TestNewDbClient(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	viper.ReadInConfig()
	db, err := NewDbClient(viper.GetStringMap("database"), nil)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate()
}
