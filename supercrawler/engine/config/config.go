package config

import (
	"time"

	"github.com/spf13/viper"
)

type CrawlerConfig struct {
	Delay       time.Duration
	Parallelism int
	UserAgent   string
}

type QueueConfig struct {
	BookQueue    string
	ChapterQueue string
	SitemapQueue string
}

type SitemapConfig struct {
	SitemapApi       string
	SitemapIndex     string
	SitemapIndexList string
	SitemapIndexLoc  string
	SitemapUrls      string
}

var (
	Queue   *QueueConfig
	Crawler *CrawlerConfig
	Sitemap *SitemapConfig
)

func Read(path string) error {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	if Crawler == nil {
		Crawler = &CrawlerConfig{
			Delay:       viper.GetDuration("crawler.delay") * time.Second,
			Parallelism: viper.GetInt("crawler.parallelism"),
			UserAgent:   viper.GetString("crawler.user_agent"),
		}
	}

	if Queue == nil {
		Queue = &QueueConfig{
			BookQueue:    viper.GetString("crawler.book_queue"),
			ChapterQueue: viper.GetString("crawler.chapter_queue"),
			SitemapQueue: viper.GetString("crawler.sitemap_queue"),
		}
	}

	if Sitemap == nil {
		Sitemap = &SitemapConfig{
			SitemapApi:       viper.GetString("crawler.sitemap_api"),
			SitemapIndex:     viper.GetString("crawler.sitemap_index"),
			SitemapIndexList: viper.GetString("crawler.sitemap_index_list"),
			SitemapIndexLoc:  viper.GetString("crawler.sitemap_index_loc"),
			SitemapUrls:      viper.GetString("crawler.sitemap_urls"),
		}
	}
	return nil
}
