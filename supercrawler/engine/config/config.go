package config

import (
	"time"

	"github.com/spf13/viper"
)

type CrawlerConfig struct {
	Delay       time.Duration
	Parallelism int
	UserAgent   string
	Queue       string
}

type QueueConfig struct {
	BookQueue    string
	ChapterQueue string
	SitemapQueue string
}

type SitemapConfig struct {
	SitemapIndexRoot string
	SitemapIndexList string
	SitemapLoc       string
	SitemapUrlRoot   string
	SitemapUrlList   string
}

var (
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
			Queue:       viper.GetString("crawler.queue"),
		}
	}

	if Sitemap == nil {
		Sitemap = &SitemapConfig{
			SitemapLoc:       viper.GetString("crawler.sitemap_loc"),
			SitemapIndexRoot: viper.GetString("crawler.sitemap_index_root"),
			SitemapIndexList: viper.GetString("crawler.sitemap_index_list"),
			SitemapUrlRoot:   viper.GetString("crawler.sitemap_url_root"),
			SitemapUrlList:   viper.GetString("crawler.sitemap_url_list"),
		}
	}
	return nil
}
