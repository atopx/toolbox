package create

import (
	"superserver/domain/public_service"
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
	Title   string                  `json:"title"`   // 标题
	Content string                  `json:"content"` // 内容
	Topics  []*public_service.Label `json:"topics"`
	Labels  []*public_service.Label `json:"labels"` // 标签
}

type Reply struct{}
