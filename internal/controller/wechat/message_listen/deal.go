package message_listen

import (
	"net/http"
	"time"
)

func (ctl *Controller) Deal() {
	reply := xml{
		ToUserName:   ctl.param.FromUserName,
		FromUserName: ctl.param.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      ctl.param.MsgType,
		Content:      ctl.param.Content,
	}
	ctl.Context.XML(http.StatusOK, reply)
}
