package models

import "gorm.io/gorm"

type Book struct {
	Base
	Name       string `json:"name"`        // 名称
	Author     string `json:"author"`      // 作者
	Label      string `json:"label"`       // 标签
	Src        string `json:"src"`         // 链接
	Cover      string `json:"cover"`       // 封面
	Intro      string `json:"intro"`       // 简介
	State      string `json:"state"`       // 状态
	LastModify int64  `json:"last_modify"` // 最后修改时间 时间戳：秒
}

func (m *Book) Connect() *gorm.DB {
	return m.conn.Table(TablenameBook)
}

func NewBookClient(db *gorm.DB) *Book {
	return &Book{Base: Base{conn: db}}
}

const (
	BookStatusNew     = "NEW"     // 新的, 未采集
	BookStatusPending = "PENDING" // 采集过的 就绪 可以更新
	BookStatusRunning = "RUNNING" // 采集中
	BookStatusFinal   = "FINAL"   // 终态 不再更新
)
