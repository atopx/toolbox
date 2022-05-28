package message_listen

import (
	"net/http"
	"time"
)

func (ctl *Controller) Deal() {
	content := "123123123"
	reply := Reply{
		ToUserName:   ctl.param.FromUserName,
		FromUserName: ctl.param.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
	ctl.NewOkResponse(http.StatusOK, reply)
}
