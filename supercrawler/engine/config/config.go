package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
)

type Crawler struct {
	Delay       time.Duration
	Parallelism int
	Threshold   int
	UserAgent   string
}

type Queue struct {
	BookQueue    string
	ChapterQueue string
}

type Cron struct {
	SitemapCron string
	CrawlerCron string
}

type Sitemap struct {
	SitemapApi       string
	SitemapIndex     string
	SitemapIndexList string
	SitemapIndexLoc  string
	SitemapUrls      string
}

type Config struct {
	Crawler
	Queue
	Cron
	Sitemap
}

func Get() *Config {
	return &Config{
		Crawler{
			Delay:       viper.GetDuration("crawler.delay") * time.Second,
			Parallelism: viper.GetInt("crawler.parallelism"),
			Threshold:   viper.GetInt("crawler.threshold"),
			UserAgent:   viper.GetString("crawler.user_agent"),
		},
		Queue{
			BookQueue:    viper.GetString("crawler.book_queue"),
			ChapterQueue: viper.GetString("crawler.chapter_queue"),
		},
		Cron{
			SitemapCron: viper.GetString("crawler.sitemap_cron"),
			CrawlerCron: viper.GetString("crawler.crawler_cron"),
		},
		Sitemap{
			SitemapApi:       viper.GetString("crawler.sitemap_api"),
			SitemapIndex:     viper.GetString("crawler.sitemap_index"),
			SitemapIndexList: viper.GetString("crawler.sitemap_index_list"),
			SitemapIndexLoc:  viper.GetString("crawler.sitemap_index_loc"),
			SitemapUrls:      viper.GetString("crawler.sitemap_urls"),
		},
	}
}

const (
	Site         = "https://www.ibiquge.info"
	Delay        = time.Second
	Parallelism  = 8
	Threshold    = 64
	BookQueue    = "novel_crawler_book"
	ChapterQueue = "novel_crawler_chapter"
)

const (
	SitemapApi       = Site + "/api/sitemap.xml"
	SitemapIndex     = "/sitemapindex"
	SitemapIndexList = "/sitemap"
	SitemapIndexLoc  = "loc"
	SitemapUrlList   = "/urlset/url"
)

const (
	SitemapCron = "0 1 * * ?"
	CrawlerCron = "*/1 * * * *"
)
