package models

import "gorm.io/gorm"

type Chapter struct {
	Base
	BookId  int    `json:"book_id"` // book-id
	Code    int    `json:"code"`    // 章节code
	Src     string `json:"src"`     // 链接
	Title   string `json:"title"`   // 章节标题
	State   int    `json:"state"`   // 状态
	Content string `json:"content"` // 章节内容
}

func (m *Chapter) Connect() *gorm.DB {
	return m.conn.Table(TablenameChapter)
}

func NewChapterClient(db *gorm.DB) *Chapter {
	return &Chapter{Base: Base{conn: db}}
}
