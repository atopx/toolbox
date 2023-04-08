package pkg

import (
	"context"
	"superserver/domain/auth_service"
	"superserver/domain/public/common"
	"superserver/domain/public_service"
	"testing"

	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
)

func TestNewGrpcClient(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	_ = viper.ReadInConfig()
	conn, err := NewGrpcClient(viper.GetStringSlice("service"))
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

func TestNewGrpcClient_PublicService(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	_ = viper.ReadInConfig()
	conn, err := NewGrpcClient(viper.GetStringSlice("service"))
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	header := metadata.New(map[string]string{"trace_id": "123123"})
	ctx = metadata.NewOutgoingContext(ctx, header)
	reply, err := public_service.NewPublicServiceClient(conn).OperateLabel(ctx, &public_service.OperateLabelParams{
		Header:  &common.Header{},
		Operate: common.Operation_OPERATION_CREATE,
		Data: &public_service.Label{
			Source: 1,
			Name:   "fpeaiohfao[we",
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("reply: %+v", reply)
}

func TestNewGrpcClient_PublicService_ListLabel(t *testing.T) {
	viper.SetConfigFile("../config.yaml")
	_ = viper.ReadInConfig()
	conn, err := NewGrpcClient(viper.GetStringSlice("service"))
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	header := metadata.New(map[string]string{"trace_id": "123123"})
	ctx = metadata.NewOutgoingContext(ctx, header)
	reply, err := public_service.NewPublicServiceClient(conn).ListLabel(ctx, &public_service.ListLabelParams{
		Header: &common.Header{},
		Pager:  &common.Pager{Index: 1, Size: 10},
		Filter: &public_service.LabelFilter{
			Ids:             []int32{1},
			Sources:         nil,
			Names:           nil,
			DeleteTimeRange: nil,
			CreateTimeRange: nil,
			UpdateTimeRange: nil,
			Creators:        nil,
			Updaters:        nil,
			Keywords:        &public_service.LabelFilter_Keywords{Name: "peaio"},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("reply: %+v", reply)
}
