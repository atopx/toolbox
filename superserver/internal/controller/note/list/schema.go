package list

import (
	"superserver/domain/note_service"
	"superserver/domain/public/common"
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
	Pager  *common.Pager  `json:"pager"`
	Sorts  []*common.Sort `json:"sorts"`
	Filter struct {
		PublicSelect common.BooleanScope `json:"public_elect"`
		Keyword      string              `json:"keyword"`
		TopicIds     []int32             `json:"topic_ids"`
		LabelIds     []int32             `json:"label_ids"`
	} `json:"filter"`
}

type Reply struct {
	Pager *common.Pager        `json:"pager"`
	List  []*note_service.Note `json:"list"`
}
