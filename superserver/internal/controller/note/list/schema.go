package list

import (
	"superserver/domain/public/common"
	"superserver/internal/controller"
	"superserver/internal/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Pager  *common.Pager  `json:"pager"`
	Sorts  []*common.Sort `json:"sorts"`
	Filter struct {
		Keyword  string  `json:"keyword"`
		TopicIds []int32 `json:"topic_ids"`
		LabelIds []int32 `json:"label_ids"`
	} `json:"filter"`
}

type Reply struct {
	Pager *common.Pager `json:"pager"`
	List  []model.Note  `json:"list"`
}
