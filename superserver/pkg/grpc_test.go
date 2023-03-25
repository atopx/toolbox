package pkg

import (
	"context"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
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
	reply, err := auth_service.NewAuthServiceClient(conn).OperateUser(ctx, &auth_service.OperateUserParams{
		Header:  &common.Header{},
		Operate: 2,
		Data: &auth_service.User{
			Id:       15,
			Username: "2a3112312422342226333333333523",
			Password: "mengfei6522323",
			Status:   auth_service.UserStatus_USER_STATUS_DISABLED,
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("reply: %+v", reply)
}
