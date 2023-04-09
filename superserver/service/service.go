package service

import (
	"google.golang.org/grpc"
	"superserver/domain/auth_service"
	"superserver/domain/note_service"
	"superserver/domain/public_service"
	"superserver/pkg"

	"github.com/spf13/viper"
)

var client *Client

type Client struct {
	conn   *grpc.ClientConn
	Public public_service.PublicServiceClient
	Auth   auth_service.AuthServiceClient
	Note   note_service.NoteServiceClient
}

func (c *Client) tryConnectService() (err error) {
	c.conn, err = pkg.NewGrpcClient(viper.GetStringSlice("service"))
	return err
}

func (c *Client) initClient() {
	if c.conn == nil {
		if err := c.tryConnectService(); err != nil {
			panic(err)
		}
	}
	c.Public = public_service.NewPublicServiceClient(c.conn)
	c.Auth = auth_service.NewAuthServiceClient(c.conn)
	c.Note = note_service.NewNoteServiceClient(c.conn)
}

func GetClient() *Client {
	if client == nil {
		client = &Client{}
		client.initClient()
	}
	return client
}
