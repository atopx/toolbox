package create

import (
	"errors"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/url"
	"strings"
	"superserver/common/interface/ecode_iface"
	"superserver/common/interface/novel_iface"
	"superserver/common/logger"
	"superserver/internal/model/novel"
	"time"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	target, err := url.Parse(params.Target)
	if err != nil {
		ctl.NewErrorResponse(ecode_iface.ECode_INVALID_PARAMS, "无效的URL地址")
		return
	}
	db := ctl.GetDatabase()
	var novelPo novel.Novel
	err = db.Model(&novelPo).Where("source = ?", target.String()).First(&novelPo).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		ctl.NewErrorResponse(ecode_iface.ECode_REQUEST_CONFLICT, "书源已存在")
		return
	}

	if err != nil {
		logger.Error(ctl.Context, "create novel first by source failed", zap.Error(err))
		ctl.NewErrorResponse(ecode_iface.ECode_DB_ERROR, "系统错误, 请联系管理员")
		return
	}

	novelPo.Source = params.Target
	novelPo.Creator = ctl.GetOperator()
	novelPo.Updater = ctl.GetOperator()

	tx := db.Begin()
	reply := &Reply{}
	var chapters []*novel.Chapter
	chapterDup := make(map[string]struct{})

	defer func() {
		if err != nil {
			tx.Rollback()
			ctl.NewErrorResponse(ecode_iface.ECode_SYSTEM_ERROR, "系统错误, 请联系管理员")
		} else {
			reply.Chapters = chapters
			reply.Novel = NovelVo{
				Id:     novelPo.Id,
				Name:   novelPo.Name,
				Source: novelPo.Source,
				Intro:  novelPo.Intro,
				Author: novelPo.Author,
				Lasted: novelPo.Lasted,
				Status: novelPo.Status.String(),
			}
			ctl.NewOkResponse(reply)
		}
	}()

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"),
		colly.IgnoreRobotsTxt(),
		colly.AllowedDomains(target.Hostname()),
	)
	c.DisableCookies()

	// 错误处理
	c.OnError(func(response *colly.Response, innerErr error) {
		logger.Error(ctl.Context, "create novel http request failed", zap.Error(innerErr), zap.String("target", params.Target))
		err = innerErr
		return
	})

	// 解析标题
	c.OnHTML("div#info", func(infoElement *colly.HTMLElement) {
		novelPo.Name = infoElement.ChildText("h1")
		authorTemp := strings.Split(infoElement.ChildText("p:nth-child(2)"), "：")
		if len(authorTemp) > 1 {
			novelPo.Author = authorTemp[len(authorTemp)-1]
		}
		lastedTemp := strings.Split(infoElement.ChildText("p:nth-child(4)"), "：")
		if len(lastedTemp) > 0 {
			lastedStr := lastedTemp[len(lastedTemp)-1]
			if t, err := time.Parse("2006-01-02 15:04:05", lastedStr); err == nil {
				novelPo.Lasted = t.Local().Unix()
			}
		}
	})
	// 解析简介
	c.OnHTML("div#intro", func(infoElement *colly.HTMLElement) {
		novelPo.Intro = strings.TrimSpace(infoElement.Text)
	})

	// 解析章节
	c.OnHTML("div#list", func(e *colly.HTMLElement) {
		e.ForEach("dd a", func(i int, element *colly.HTMLElement) {
			source := element.Request.AbsoluteURL(element.Attr("href"))
			if _, ok := chapterDup[source]; !ok && source != "" {
				chapterDup[source] = struct{}{}
				chapters = append(chapters, &novel.Chapter{
					Name:   element.Text,
					Source: source,
					Status: novel_iface.ScanStatus_SCAN_PENDING,
				})
			}
		})
	})
	if err = c.Visit(target.String()); err != nil {
		logger.Error(ctl.Context, "create novel visit failed", zap.Error(err), zap.String("target", params.Target))
		return
	}
	novelPo.Status = novel_iface.NovelStatus_NOVEL_ENABLED
	if err = tx.Create(&novelPo).Error; err != nil {
		logger.Error(ctl.Context, "create novel tx.Create failed", zap.Error(err))
		return
	}

	for _, chapter := range chapters {
		chapter.NovelId = novelPo.Id
	}
	tx.Commit()
	if err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&chapters).Error; err != nil {
		logger.Error(ctl.Context, "create novel chapter db.Create failed", zap.Error(err))
		return
	}
}
