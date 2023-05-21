package deltag

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	NoteId  int32 `json:"note_id"`
	LabelId int32 `json:"label_id"`
	IsTopic bool  `json:"is_topic"`
}

type Reply struct{}
