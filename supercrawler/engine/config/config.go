package config

import "time"

const (
	UserAgent = "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"
)

const (
	Site        = "https://www.ibiquge.info"
	Delay       = time.Second
	Parallelism = 8
	QueueName   = "novel_crawler"
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
	BookCron    = "*/20 * * * *"
)
