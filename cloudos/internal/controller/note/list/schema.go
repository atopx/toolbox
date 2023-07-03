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
	Pager           *pb.Pager    `json:"pager"`
	Topic           string       `json:"topic"`
	Label           string       `json:"label"`   // like labels
	Keyword         string       `json:"keyword"` // like name or content
	UpdateTimeRange *pb.I64Range `json:"updateTimeRange"`
}

type Reply struct {
	Pager *pb.Pager `json:"pager"`
	List  []Item    `json:"list"`
}

type Item struct {
	Id         int64  `json:"id"`
	Title      string `json:"title"`
	Topic      string `json:"topic"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
