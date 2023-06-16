package list

import (
	"cloudos/common/pb"
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
	Pager    *pb.Pager `json:"pager"`
	FolderId int64     `json:"parentId"`
	TopicId  int64     `json:"topicId"`
	Label    string    `json:"label"`
	Keyword  string    `json:"keyword"`
}

type Reply struct {
	Pager *pb.Pager `json:"pager"`
	List  []Item    `json:"list"`
}

type Item struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	TopicName  string  `json:"topicName"`
	CreateTime string  `json:"createTime"`
	UpdateTime string  `json:"updateTime"`
	Labels     []Label `json:"labels"`
}

type Label struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
