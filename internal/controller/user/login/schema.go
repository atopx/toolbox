package login

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Reply struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
