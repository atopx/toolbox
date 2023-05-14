package sitemap

import (
	"github.com/antchfx/xmlquery"
	"github.com/gocolly/colly/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	crawler.ctor.OnXML(config.Sitemap.SitemapIndexRoot, crawler.sitemapIndex)
	crawler.ctor.OnXML(config.Sitemap.SitemapUrlRoot, crawler.sitemapUrlList)
	return &crawler
}

func (c *Crawler) sitemapIndex(element *colly.XMLElement) {
	for _, node := range element.DOM.(*xmlquery.Node).SelectElements(config.Sitemap.SitemapIndexList) {
		_ = element.Request.Visit(node.SelectElement(config.Sitemap.SitemapLoc).InnerText())
	}
}

func (c *Crawler) sitemapUrlList(element *colly.XMLElement) {
	books := make([]models.Book, 0, 500)
	client := models.NewBookClient(c.db)
	for _, node := range element.DOM.(*xmlquery.Node).SelectElements(config.Sitemap.SitemapUrlList) {
		books = append(books, models.Book{
			Src:   node.SelectElement(config.Sitemap.SitemapLoc).InnerText(),
			State: models.PENDING,
		})
	}
	client.Connect().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "src"}},
		DoNothing: true,
	}).Create(books)
}

func (c *Crawler) request(request *colly.Request) {
	logger.Debug("sitemap crawler request", zap.String("src", request.URL.String()))
}
func (c *Crawler) response(response *colly.Response) {
	logger.Debug("sitemap crawler response", zap.Int("status_code", response.StatusCode), zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) scraped(response *colly.Response) {
	logger.Info("sitemap crawler success", zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) error(response *colly.Response, err error) {
	logger.Error("sitemap crawler failed", zap.Error(err), zap.String("src", response.Request.URL.String()))
}

func (c *Crawler) Wait() {
	c.ctor.Wait()
}

func (c *Crawler) Visit(url string) error {
	return c.ctor.Visit(url)
}
