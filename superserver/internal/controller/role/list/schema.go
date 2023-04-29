package list

import (
	"github.com/gin-gonic/gin"
	"superserver/domain/public/common"
	"superserver/internal/controller"
	"superserver/internal/model"
)

type Controller struct {
	*controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{controller.New(ctx, new(Params))}
}

type Params struct {
	Pager  *common.Pager `json:"pager"`
	Filter struct {
		Keyword string `json:"keyword"`
	} `json:"filter"`
}

type Reply struct {
	Pager *common.Pager `json:"pager"`
	List  []model.Role  `json:"list"`
}
