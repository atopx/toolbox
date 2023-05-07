package novel

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Chapter struct {
	Title   string
	Content string
}

type Book struct {
	start    time.Time
	Name     string
	Author   string
	Chapters []*Chapter
}

type Crawler struct {
	collector *colly.Collector
	replacer  *strings.Replacer
}

func New() *Crawler {
	crawler := &Crawler{
		collector: colly.NewCollector(),
		replacer:  strings.NewReplacer("\u00a0\u00a0", "\n"),
	}
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
	//log.Printf("request: %s", request.URL)
}
func (c *Crawler) response(response *colly.Response) {
	//fmt.Println("OnResponse")
}
func (c *Crawler) scraped(response *colly.Response) {
	//log.Printf("success: %s", response.Request.URL)
}
func (c *Crawler) error(response *colly.Response, err error) {
	log.Printf("crawler %s failed: %s", response.Request.URL, err.Error())
}

const (
	BookInfoSelector       = "#info"
	ChapterListSelector    = "#list"
	ChapterLinkSelector    = "dd:nth-child(n+14) a"
	ChapterContentSelector = "#content"
)

func (c *Crawler) Save(book Book) error {
	file, err := os.OpenFile(fmt.Sprintf("%s-%s.txt", book.Name, book.Author), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	for _, chapter := range book.Chapters {
		_, _ = file.WriteString(chapter.Title)
		_, _ = file.WriteString("\n")
		_, _ = file.WriteString(chapter.Content)
		_, _ = file.WriteString("\n")
	}
	log.Printf("save success, total cost %s", time.Since(book.start))
	return file.Close()
}

func (c *Crawler) Do(url string) error {
	book := Book{start: time.Now().Local()}
	records := make(map[string]int)

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
		c.collector.OnHTMLDetach(BookInfoSelector)
	})

	c.collector.OnHTML(ChapterListSelector, func(element *colly.HTMLElement) {

		element.ForEach(ChapterLinkSelector, func(i int, element *colly.HTMLElement) {
			link := element.Request.AbsoluteURL(element.Attr("href"))
			records[link] = i
			chapter := &Chapter{Title: element.Text}
			book.Chapters = append(book.Chapters, chapter)
			//log.Printf("start crawler %s: %s", chapter.Title, link)
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
	log.Printf("crawled %d page, cost %s", len(book.Chapters)+1, time.Since(book.start))
	return c.Save(book)
}
