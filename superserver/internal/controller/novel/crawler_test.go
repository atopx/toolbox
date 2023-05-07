package novel

import (
	"testing"
)

func TestCrawler_Do(t *testing.T) {

	addr := "https://www.ibiquge.info/131_131180/"

	crawler := New()
	if err := crawler.Do(addr); err != nil {
		t.Fatal(err)
	}
}
