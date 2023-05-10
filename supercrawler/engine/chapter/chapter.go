package chapter

import (
	"net/http"
	"strings"
	"supercrawler/common/logger"
	"supercrawler/engine/config"
	"supercrawler/models"
	"time"

	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
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
	crawler.ctor.UserAgent = config.Crawler.UserAgent
	crawler.ctor.DisableCookies()
	crawler.ctor.WithTransport(&http.Transport{DisableKeepAlives: true})
	_ = crawler.ctor.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: config.Crawler.Delay,
		Parallelism: config.Crawler.Parallelism,
	})

	crawler.ctor.OnRequest(crawler.request)
	crawler.ctor.OnError(crawler.error)
	crawler.ctor.OnScraped(crawler.scraped)
	crawler.ctor.OnResponse(crawler.response)
	crawler.ctor.OnHTML("#main", crawler.chapter)
	return &crawler
}

func (c *Crawler) request(request *colly.Request) {
	logger.Debug("chapter crawler request", zap.String("src", request.URL.String()))
}
func (c *Crawler) response(response *colly.Response) {
	logger.Debug("chapter crawler response", zap.Int("status_code", response.StatusCode), zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) scraped(response *colly.Response) {
	logger.Info("chapter crawler success", zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) error(response *colly.Response, err error) {
	logger.Error("chapter crawler failed", zap.Error(err), zap.String("src", response.Request.URL.String()))
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

func (c *Crawler) Wait() {
	c.ctor.Wait()
}

func (c *Crawler) Visit(url string) error {
	return c.ctor.Visit(url)
}
