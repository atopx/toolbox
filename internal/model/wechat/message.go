package wechat

import (
	"fmt"
	"log"
	"time"
)

type Message struct {
	ToUserName   string `json:"ToUserName" xml:"ToUserName"`
	FromUserName string `json:"FromUserName" xml:"FromUserName"`
	CreateTime   int64  `json:"CreateTime" xml:"CreateTime"`
	MsgType      string `json:"MsgType" xml:"MsgType"`
	Content      string `json:"Content" xml:"Content"`
	MsgId        string `json:"MsgId" xml:"MsgId"`
	MsgDataId    string `json:"MsgDataId" xml:"MsgDataId"`
	Idx          string `json:"Idx" xml:"Idx"`
}

func (msg *Message) ReplyText(content string) Message {
	fmt.Println(*msg)
	log.Printf("receive: %s", msg.Content)
	log.Printf("reply: %s", content)
	reply := Message{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}
	fmt.Println(reply)
	return reply
}
