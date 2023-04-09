package info

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
	NoteId int32 `from:"note_id"`
}

type Reply struct {
	Note   *note_service.Note      `json:"note"`
	Topic  *note_service.NoteTopic `json:"topic"`
	Labels []*public_service.Label `json:"labels"`
}
