package pkg

import (
	"context"
	"superserver/domain/auth_service"
	"testing"

	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

func TestNewGrpcClient(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	viper.ReadInConfig()
	conn, err := NewGrpcClient(viper.GetStringSlice("service.auth"))
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	header := metadata.New(map[string]string{"trace_id": "123123"})
	ctx = metadata.NewOutgoingContext(ctx, header)
	_, err = auth_service.NewAuthServiceClient(conn).Authorization(ctx, &auth_service.AuthorizationParams{})
	if err != nil {
		t.Error(err)
	}
}
