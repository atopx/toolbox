package message_listen

import (
	"net/http"
	"time"
	"toolbox/common/logger"

	"go.uber.org/zap"
)

func (ctl *Controller) Deal() {
	process := NewHandler(ctl.param.Content)
	reply := xml{
		ToUserName:   ctl.param.FromUserName,
		FromUserName: ctl.param.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      ctl.param.MsgType,
		Content:      process.Deal(),
	}
	logger.Debug(ctl.Context, "message_listen", zap.String("content", ctl.param.Content))
	ctl.Context.XML(http.StatusOK, reply)
}
