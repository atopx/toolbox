package export

import (
	"bytes"
	"superserver/internal/model/novel"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	db := ctl.GetDatabase()

	var chapters []novel.Chapter
	db.Model(&novel.Chapter{}).Where("novel_id=?", params.Id).
		Select("name, content").
		Order("source").Find(&chapters)

	var buf bytes.Buffer

	for _, chapter := range chapters {
		buf.WriteString(chapter.Name)
		buf.WriteString("\n\n")
		buf.WriteString(chapter.Content)
	}
	// 	c.Header("content-disposition", "attachment; filename="+template.Filename)
	//	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", template.ToBytes())
	buf.Bytes()
}
