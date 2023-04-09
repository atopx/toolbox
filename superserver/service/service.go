package service

import (
	"superserver/domain/auth_service"
	"superserver/domain/public_service"
	"superserver/pkg"

	"github.com/spf13/viper"
)

var client *Client

type Client struct {
	Auth   auth_service.AuthServiceClient
	Public public_service.PublicServiceClient
}

func initClient() {
	grpcClient, err := pkg.NewGrpcClient(viper.GetStringSlice("service"))
	if err != nil {
		panic(err)
	}
	client = &Client{
		Auth:   auth_service.NewAuthServiceClient(grpcClient),
		Public: public_service.NewPublicServiceClient(grpcClient),
	}
}

func GetClient() *Client {
	if client == nil {
		initClient()
	}
	return client
}
