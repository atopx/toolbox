package novels

import (
	"strings"
	"superserver/common/interface/novel_iface"
	"superserver/internal/scheduler/base"
	"time"

	"github.com/gocolly/colly/v2"
	"gorm.io/gorm"
)

type NovelScheduler struct {
	base.Scheduler
	collector *colly.Collector
}

func New(db *gorm.DB) *NovelScheduler {
	s := &NovelScheduler{Scheduler: base.NewScheduler(db)}
	s.initCollector()
	return s
}

func (s *NovelScheduler) initCollector() {
	s.collector = colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)"),
		colly.IgnoreRobotsTxt(),
	)
	s.collector.DisableCookies()
	s.collector.Async = true
	s.collector.Limit(&colly.LimitRule{
		RandomDelay: time.Second,
		Parallelism: 4,
	})

	// 请求开始
	s.collector.OnRequest(func(request *colly.Request) {
		s.Db.Table("su_novel_chapter").
			Where("source = ?", request.URL.String()).
			Update("status", novel_iface.ScanStatus_SCAN_RUNNING)
	})

	// 错误处理
	s.collector.OnError(func(response *colly.Response, err error) {
		s.Db.Table("su_novel_chapter").
			Where("source = ?", response.Request.URL.String()).
			Updates(map[string]interface{}{
				"status":      novel_iface.ScanStatus_SCAN_ERROR,
				"message":     err.Error(),
				"update_time": time.Now().Local().Unix(),
			})
	})

	// 解析内容
	s.collector.OnHTML("div#content", func(element *colly.HTMLElement) {
		text := strings.TrimSpace(element.Text)
		s.Db.Table("su_novel_chapter").
			Where("source = ?", element.Request.URL.String()).
			Updates(map[string]interface{}{
				"status":      novel_iface.ScanStatus_SCAN_SUCCESS,
				"content":     text,
				"update_time": time.Now().Local().Unix(),
			})
	})
}

func (s *NovelScheduler) Cron() string {
	// 每天凌晨5点更新
	return "0 5 * * *"
}

func (s *NovelScheduler) Start() {
	var targets []string
	s.Db.Table("su_novel_chapter").Select("source").Where("status = ?", novel_iface.ScanStatus_SCAN_PENDING).
		Order("id").Find(&targets)
	for _, target := range targets {
		s.collector.Visit(target)
	}

	s.collector.Wait()
}
