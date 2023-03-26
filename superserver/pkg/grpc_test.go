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
	reply, err := auth_service.NewAuthServiceClient(conn).ListUser(ctx, &auth_service.ListUserParams{
		Header: &common.Header{},
		Pager: &common.Pager{
			Index:    1,
			Size:     20,
			Count:    0,
			Disabled: true,
		},
		Sorts:  []*common.Sort{{Field: "create_time", Direction: common.SortDirection_SORT_ASC}},
		Filter: &auth_service.UserFilter{States: []auth_service.UserStatus{auth_service.UserStatus_USER_STATUS_DISABLED}},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("reply: %+v", reply)
}
