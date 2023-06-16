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
	ParentId int64     `json:"parentId"`
	Domain   string    `json:"domain"`
	Keyword  string    `json:"keyword"`
}

type Reply struct {
	Pager *pb.Pager `json:"pager"`
	List  []Folder  `json:"list"`
}

type Folder struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}
