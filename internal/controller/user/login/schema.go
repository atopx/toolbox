package login

import (
	"superserver/internal/controller"
	"superserver/internal/model"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	param *Param
	controller.Controller
}

func NewController(ctx *gin.Context) (*Controller, error) {
	ctl := Controller{param: new(Param)}
	ctl.ContextLoader(ctx)
	if err := ctl.NewRequestParam(ctl.param); err != nil {
		return nil, err
	}
	return &ctl, nil
}

type Param struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Reply struct {
	Id              int64      `json:"id"`
	Username        string     `json:"username"`
	Role            model.Enum `json:"role"`
	Status          model.Enum `json:"status"`
	AvatarUrl       string     `json:"avatar_url"`
	TotalSizeLimit  int64      `json:"total_size_limit"`
	SingleSizeLimit int64      `json:"single_size_limit"`
	SizeUsed        int64      `json:"size_used"`
	LastLoginIp     string     `json:"last_login_ip"`
	LastLoginTime   int64      `json:"last_login_time"`
	CreateTime      int64      `json:"create_time"`
	UpdateTime      int64      `json:"update_time"`
}
