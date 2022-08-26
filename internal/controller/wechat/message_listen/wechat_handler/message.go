package wechat_handler

import "strings"

type MessageHandler struct{}

func (h *MessageHandler) Deal() string {
	var messages = []string{
		"ts: 时间戳工具(timestamp)",
		"wt: 天气预报(weather) TODO",
		"json: JSO对象转义、格式化",
		"wol: 远程开机(wake on lan) TODO",
		"trans: 翻译(translate) TODO",
		"compress: 字符串压缩, 去除空格",
	}
	return strings.Join(messages, "\n")
}
