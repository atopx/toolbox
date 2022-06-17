package wechat

type Message struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`     // 开发者微信号
	FromUserName string `json:"FromUserName" xml:"FromUserName"` // 发送方帐号（一个OpenID）
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`     // 消息创建时间 （整型）
	MsgType      string `json:"MsgType" xml:"MsgType"`           // 消息类型，文本为text
	Content      string `json:"Content" xml:"Content"`           // 文本消息内容
	MsgId        string `json:"MsgId" xml:"MsgId"`               // 消息id，64位整型
}
