package delete

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
	AccessId int `json:"access_id"`
	RoleId   int `json:"role_id"`
}

type Reply struct{}
