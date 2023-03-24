package service

import (
	"superserver/domain/auth_service"
	"superserver/pkg"

	"github.com/spf13/viper"
)

var client *Client

type Client struct {
	Auth auth_service.AuthServiceClient
}

func initClient() {
	authClient, err := pkg.NewGrpcClient(viper.GetStringSlice("service.auth"))
	if err != nil {
		panic(err)
	}
	client = &Client{
		Auth: auth_service.NewAuthServiceClient(authClient),
	}
}

func GetClient() *Client {
	if client == nil {
		initClient()
	}
	return client
}
