package create

import (
	"superserver/domain/note_service"
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
	Title    string                  `json:"title"`     // 标题
	Content  string                  `json:"content"`   // 内容
	FolderId int32                   `json:"folder_id"` // 文件夹ID
	Public   bool                    `json:"public"`    // 是否公开
	Topic    *note_service.NoteTopic `json:"topic"`     // 主题
	Labels   []*public_service.Label `json:"labels"`    // 标签
}

type Reply struct{}
