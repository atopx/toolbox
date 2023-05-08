package novel

import (
	"context"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"superserver/common/consts"
	"superserver/common/logger"
	"superserver/common/utils"
	"superserver/domain/public/common"
	"superserver/pkg"
	"syscall"
	"time"
)

type Chapter struct {
	Title   string
	Content string
}

type Book struct {
	src      string
	start    time.Time
	Name     string
	Author   string
	Chapters []*Chapter
}

type Crawler struct {
	collector *colly.Collector
	replacer  *strings.Replacer
	store     *redis.Client
	ctx       context.Context
	cancel    context.CancelFunc
	queue     *redis.PubSub
}

func New() *Crawler {
	crawler := &Crawler{
		collector: colly.NewCollector(),
		replacer:  strings.NewReplacer("\u00a0\u00a0", "\n"),
		store:     pkg.NewRedisClient(viper.GetStringMap("redis")),
	}

	crawler.ctx, crawler.cancel = context.WithCancel(context.Background())
	crawler.ctx = context.WithValue(crawler.ctx, consts.RequestHeader, &common.Header{
		TraceId: utils.NewTraceId(),
		Source:  "novel-crawler",
	})
	crawler.queue = crawler.store.Subscribe(crawler.ctx, "crawler-novel")
	crawler.collector.Async = true
	crawler.collector.TraceHTTP = true
	crawler.collector.DetectCharset = true
	crawler.collector.UserAgent = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
	crawler.collector.DisableCookies()
	crawler.collector.IgnoreRobotsTxt = true
	crawler.collector.WithTransport(&http.Transport{DisableKeepAlives: true})
	crawler.collector.OnRequest(crawler.request)
	crawler.collector.OnError(crawler.error)
	crawler.collector.OnScraped(crawler.scraped)
	crawler.collector.OnResponse(crawler.response)
	return crawler
}

func (c *Crawler) request(request *colly.Request) {
	logger.Debug(c.ctx, "request", zap.String("src", request.URL.String()))
}
func (c *Crawler) response(response *colly.Response) {
	logger.Debug(c.ctx, "response", zap.Int("status_code", response.StatusCode), zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) scraped(response *colly.Response) {
	logger.Debug(c.ctx, "success", zap.String("src", response.Request.URL.String()))
}
func (c *Crawler) error(response *colly.Response, err error) {
	logger.Error(c.ctx, "failed", zap.Error(err), zap.String("src", response.Request.URL.String()))
}

const (
	BookInfoSelector       = "#info"
	ChapterListSelector    = "#list"
	ChapterLinkSelector    = "dd:nth-child(n+14) a"
	ChapterContentSelector = "#content"
)

func (c *Crawler) Save(book Book) error {
	buf := new(strings.Builder)
	for _, chapter := range book.Chapters {
		buf.WriteString(chapter.Title)
		buf.WriteString("\n")
		buf.WriteString(chapter.Content)
		buf.WriteString("\n")
	}
	c.store.HSet(
		c.ctx,
		fmt.Sprintf("novel::%s", book.src),
		"name", book.Name,
		"author", book.Author,
		"content", buf.String(),
	)
	logger.Info(c.ctx, "saved", zap.String("src", book.src), zap.Duration("cost", time.Since(book.start)))
	return nil
}

func (c *Crawler) scheduler(url string) error {
	book := Book{start: time.Now().Local(), src: url}
	records := make(map[string]int)
	header := c.ctx.Value(consts.RequestHeader).(*common.Header)
	header.TraceId = utils.NewTraceId()
	header.Source = url

	if err := c.collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: time.Second,
		Parallelism: 8,
	}); err != nil {
		return err
	}

	c.collector.OnHTML(BookInfoSelector, func(element *colly.HTMLElement) {
		book.Name = element.ChildText("h1")
		book.Author = element.ChildTexts("a")[0]
		logger.Debug(c.ctx, "crawler", zap.Any("book", book), zap.String("src", url))
		c.collector.OnHTMLDetach(BookInfoSelector)
	})

	c.collector.OnHTML(ChapterListSelector, func(element *colly.HTMLElement) {

		element.ForEach(ChapterLinkSelector, func(i int, element *colly.HTMLElement) {
			link := element.Request.AbsoluteURL(element.Attr("href"))
			records[link] = i
			chapter := &Chapter{Title: element.Text}
			book.Chapters = append(book.Chapters, chapter)
			logger.Debug(c.ctx, "visit", zap.String("title", chapter.Title), zap.String("src", link))
			_ = element.Request.Visit(link)
		})
		c.collector.OnHTMLDetach(ChapterListSelector)

		c.collector.OnHTML(ChapterContentSelector, func(element *colly.HTMLElement) {
			content := element.DOM.After("p").Before("div").Contents().Text()
			book.Chapters[records[element.Request.URL.String()]].Content = c.replacer.Replace(strings.TrimSpace(content))
		})
	})

	if err := c.collector.Visit(url); err != nil {
		return err
	}
	c.collector.Wait()
	logger.Debug(c.ctx, "crawled", zap.Int("pages", len(book.Chapters)), zap.Duration("cost", time.Since(book.start)), zap.String("src", url))
	return c.Save(book)
}

func (c *Crawler) Listen() {
	go func() {
		for message := range c.queue.Channel(redis.WithChannelSize(1), redis.WithChannelHealthCheckInterval(time.Minute)) {
			if err := c.scheduler(message.Payload); err != nil {
				logger.Error(c.ctx, "failed", zap.String("src", message.Payload), zap.Error(err))
			}
		}
	}()
	log.Println("start novel crawler")
	manager := make(chan os.Signal, 1)
	signal.Notify(manager, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-manager:
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			c.Stop()
		}
	}
}

func (c *Crawler) Stop() {
	_ = c.queue.Close()
	c.cancel()
}
