package service

import (
	"superserver/domain/auth_service"
	"superserver/domain/note_service"
	"superserver/domain/public_service"
	"superserver/pkg"

	"github.com/spf13/viper"
)

var client *Client

type Client struct {
	Public public_service.PublicServiceClient
	Auth   auth_service.AuthServiceClient
	Note   note_service.NoteServiceClient
}

func initClient() {
	grpcClient, err := pkg.NewGrpcClient(viper.GetStringSlice("service"))
	if err != nil {
		panic(err)
	}
	client = &Client{
		Public: public_service.NewPublicServiceClient(grpcClient),
		Auth:   auth_service.NewAuthServiceClient(grpcClient),
		Note:   note_service.NewNoteServiceClient(grpcClient),
	}
}

func GetClient() *Client {
	if client == nil {
		initClient()
	}
	return client
}
