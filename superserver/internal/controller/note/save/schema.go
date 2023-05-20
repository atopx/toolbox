package save

import (
	"superserver/internal/controller"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Id      int32   `json:"id"`
	Title   string  `json:"title"`   // 标题
	Content string  `json:"content"` // 内容
	Topics  []int32 `json:"topics"`
	Labels  []int32 `json:"labels"` // 标签
}

type Reply struct{}
