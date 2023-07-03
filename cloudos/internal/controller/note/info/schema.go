package info

import (
	"cloudos/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Id int64 `form:"id"`
}

type Reply struct {
	Id         int64    `json:"id"`
	Title      string   `json:"title"`
	Topic      string   `json:"topic"`
	Labels     []string `json:"labels"`
	Content    string   `json:"content"`
	CreateTime string   `json:"createTime"`
	UpdateTime string   `json:"updateTime"`
}
