package sitemap

import (
	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"supercrawler/common/logger"
	"supercrawler/engine/config"
	"supercrawler/models"
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
	crawler.ctor.OnXML(config.SitemapIndex, crawler.sitemapIndex)
	crawler.ctor.OnXML(config.SitemapUrlList, crawler.sitemapUrlList)
	return &crawler
}

func (c *Crawler) sitemapIndex(element *colly.XMLElement) {
	dom := element.DOM.(*xmlquery.Node)
	for _, node := range dom.SelectElements(config.SitemapIndexList) {
		_ = element.Request.Visit(node.SelectElement(config.SitemapIndexLoc).InnerText())
	}
}

func (c *Crawler) sitemapUrlList(element *colly.XMLElement) {
	book := models.NewBookClient(c.db)
	book.Src = element.ChildText(config.SitemapIndexLoc)
	book.State = models.BookStatusNew
	book.Connect().Where("src=?", book.Src).FirstOrCreate(&book)
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

func (c *Crawler) Start() {
	_ = c.ctor.Visit(config.SitemapApi)
	c.ctor.Wait()
}
