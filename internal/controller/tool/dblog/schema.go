package dblog

import (
	"time"
	"toolbox/internal/controller"

	"github.com/gin-gonic/gin"
)

//type Param struct {
//	JsonStr string `form:"json_str" json:"json_str"` // json字符串
//}

type Param struct {
	Level     string    `json:"level"`
	Time      time.Time `json:"time"`
	Linenum   string    `json:"linenum"`
	Func      string    `json:"func"`
	Message   string    `json:"message"`
	Svr       string    `json:"svr"`
	Traceid   string    `json:"traceid"`
	Spanid    string    `json:"spanid"`
	Trackid   int64     `json:"trackid"`
	Infc      string    `json:"infc"`
	Requester string    `json:"requester"`
	BfoUid    int       `json:"bfoUid"`
	Uid       string    `json:"uid"`
	Cost      int       `json:"cost"`
}

type Reply string

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
