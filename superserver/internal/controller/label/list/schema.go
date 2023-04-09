package list

import (
	"github.com/gin-gonic/gin"
	"superserver/domain/public/common"
	"superserver/domain/public_service"
	"superserver/internal/controller"
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
		Ids     []int32 `json:"ids"`
		Sources []int32 `json:"source"`
		Keyword string  `json:"keyword"`
	} `json:"filter"`
}

type Reply struct {
	Pager *common.Pager           `json:"pager"`
	List  []*public_service.Label `json:"list"`
}
