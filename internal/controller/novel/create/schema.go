package create

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
	"superserver/internal/model/novel"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Target string `json:"target"`
}

type Reply struct {
	Novel    NovelVo               `json:"novel"`
	Chapters []*novel.NovelChapter `json:"chapters"`
}

type NovelVo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`   // 小说书名
	Source string `json:"source"` // 目录页
	Intro  string `json:"intro"`  // 简介
	Author string `json:"author"` // 作者
	Lasted int64  `json:"lasted"` // 最后更新时间
	Status string `json:"status"` // 小说状态

}

type NovelChapterVo struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}
