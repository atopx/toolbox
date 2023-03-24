package info

import (
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct{}

type Reply struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	LoginTime  int64  `json:"login_time"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Roles      []Role `json:"roles"`
}

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
