package book

import (
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"strings"
	"supercrawler/common/logger"
	"supercrawler/common/utils"
	"supercrawler/engine/config"
	"supercrawler/models"
	"time"
)

type Crawler struct {
	db   *gorm.DB
	ctor *colly.Collector
}

func New(db *gorm.DB) *Crawler {
	crawler := Crawler{
		db:   db,
		ctor: colly.NewCollector(),
	}
	crawler.ctor.Async = true
	crawler.ctor.TraceHTTP = true
	crawler.ctor.DetectCharset = true
	crawler.ctor.IgnoreRobotsTxt = true
	crawler.ctor.AllowURLRevisit = true
	crawler.ctor.UserAgent = config.UserAgent
	crawler.ctor.DisableCookies()
	crawler.ctor.WithTransport(&http.Transport{DisableKeepAlives: true})
	_ = crawler.ctor.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: config.Delay,
		Parallelism: config.Parallelism,
	})

	crawler.ctor.OnRequest(crawler.request)
	crawler.ctor.OnError(crawler.error)
	crawler.ctor.OnScraped(crawler.scraped)
	crawler.ctor.OnResponse(crawler.response)
	crawler.ctor.OnHTML("#main", crawler.book)
	return &crawler
}

func (c *Crawler) request(request *colly.Request) {
	logger.Debug("request", zap.String("src", request.URL.String()))
}
func (c *Crawler) response(response *colly.Response) {
	logger.Debug("response", zap.Int("status_code", response.StatusCode), zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) scraped(response *colly.Response) {
	logger.Debug("success", zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) error(response *colly.Response, err error) {
	logger.Error("failed", zap.Error(err), zap.String("src", response.Request.URL.String()))
}

func (c *Crawler) book(element *colly.HTMLElement) {
	if s := element.DOM.Find("#bookdetail"); len(s.Nodes) > 0 {
		book := models.NewBookClient(c.db)
		book.Connect().Where("src=?", element.Request.URL.String()).First(&book)
		if book.State != models.BookStatusFinal {
			if err := c.parseBook(book, element); err != nil {
				logger.Error("parse book error", zap.String("src", element.Request.URL.String()), zap.Error(err))
			} else {
				if err = c.parseChapters(book, element); err != nil {
					logger.Error("parse chapters error", zap.String("src", element.Request.URL.String()), zap.Error(err))
				}
			}
		}
	} else {
		c.chapter(element)
	}
}

func (c *Crawler) chapter(element *colly.HTMLElement) {
	content := element.DOM.Find("#content").After("p").Before("div").Contents().Text()
	content = strings.ReplaceAll(strings.TrimSpace(content), "\u00a0\u00a0", "\n")
	models.NewChapterClient(c.db).Connect().Where("src = ?", element.Request.URL.String()).Updates(map[string]interface{}{
		"content":     content,
		"state":       models.BookStatusFinal,
		"update_time": time.Now().Local().Unix(),
	})
}

func (c *Crawler) parseBook(book *models.Book, element *colly.HTMLElement) error {
	book.Name = element.ChildText("#info h1")
	book.Src = element.Request.URL.String()
	book.Author = element.ChildText("#info p:nth-child(2) a")
	book.State = models.BookStatusPending
	book.Intro = strings.ReplaceAll(element.ChildText("#intro"), "\u00a0\u00a0", "\n")
	if len(book.Intro) > 4095 {
		book.Intro = book.Intro[:4095]
	}
	book.LastModify = utils.TimeLoad(element.ChildText("#info p:nth-child(4)"), "最后更新：2006-01-02 15:04:05")
	book.Cover = element.Request.AbsoluteURL(element.ChildAttr("#fmimg img", "src"))
	return book.Update(book)
}

func (c *Crawler) parseChapters(book *models.Book, element *colly.HTMLElement) error {
	var chapters []*models.Chapter
	element.ForEach("#list dl dd:nth-child(n+14) a", func(i int, element *colly.HTMLElement) {
		chapter := models.Chapter{
			BookId: book.Id,
			Code:   i + 1,
			Src:    element.Request.AbsoluteURL(element.Attr("href")),
			Title:  element.Text,
			State:  models.BookStatusPending,
		}
		chapters = append(chapters, &chapter)
		if err := models.NewChapterClient(c.db).Connect().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "src"}},
			DoUpdates: clause.AssignmentColumns([]string{"code", "title", "state", "update_time"}),
		}).Create(&chapter).Error; err != nil {
			logger.Error("parse chapter error", zap.Error(err), zap.String("src", chapter.Src))
		}
		_ = c.ctor.Visit(chapter.Src)
	})
	return nil
}

func (c *Crawler) Wait() {
	c.ctor.Wait()
}

func (c *Crawler) Visit(url string) error {
	return c.ctor.Visit(url)
}
