package novel

import (
	"superserver/common/interface/novel_iface"
	"superserver/internal/model"
)

type Novel struct {
	model.Base

	Name        string                  // 小说书名
	Source      string                  // 目录页
	Intro       string                  // 简介
	Author      string                  // 作者
	Lasted      int64                   // 最后更新时间
	Status      novel_iface.NovelStatus // 小说状态
	ScanNum     int                     // 失败次数
	NoChangeNum int                     // 未发生变化次数
	ErrorNum    int                     // 扫描失败次数
	Creator     int                     // 创建人
	Updater     int                     // 更新人
}

func (*Novel) TableName() string {
	return "su_novel"
}

type Chapter struct {
	model.Base
	NovelId int                    `json:"novel_id"`                  // 小说ID
	Name    string                 `json:"name"`                      // 章节标题
	Source  string                 `json:"source"`                    // 章节源
	Status  novel_iface.ScanStatus `json:"status"`                    // 状态
	Content string                 `json:"content" gorm:"default:''"` // 章节内容
	Message string                 `json:"message"`                   // 消息
}

func (*Chapter) TableName() string {
	return "su_novel_chapter"
}

type Task struct {
	model.Base
	NovelId int                    // 小说ID
	Status  novel_iface.ScanStatus // 状态
}
