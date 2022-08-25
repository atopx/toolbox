package gfwrules

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/pkg/errors"
)

// Cron 每天凌晨7点更新
const Cron = "0 7 * * *"

const (
	// 直连域名列表 direct.txt
	direct = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/direct.txt"
	// 代理域名列表 proxy.txt
	proxy = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/proxy.txt"
	// 广告域名列表 reject.txt
	reject = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/reject.txt"
	// 私有网络专用域名列表 private.txt
	private = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/private.txt"
	// Apple 在中国大陆可直连的域名列表 apple.txt
	apple = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/apple.txt"
	// iCloud 域名列表 icloud.txt
	icloud = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/icloud.txt"
	// Telegram 使用的 IP 地址列表 telegramcidr.txt
	telegramcidr = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/telegramcidr.txt"
	// 局域网 IP 及保留 IP 地址列表 lancidr.txt
	lancidr = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/lancidr.txt"
	// 中国大陆 IP 地址列表 cncidr.txt
	cncidr = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/cncidr.txt"
	// 需要直连的常见软件列表 applications.txt
	applications = "https://raw.githubusercontent.com/Loyalsoldier/clash-rules/release/applications.txt"
)

var targets = []string{
	direct,
	proxy,
	reject,
	private,
	apple,
	icloud,
	telegramcidr,
	lancidr,
	cncidr,
	applications,
}

func GFWDownloader() {
	for i := 0; i < len(targets); i++ {
		if err := fatch(targets[i]); err != nil {
			log.Println(err.Error())
		}
	}
}

func fatch(url string) error {
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	client := http.Client{Timeout: 300 * time.Second}
	log.Printf("[GFWDownloader] start download %s\n", url)
	response, err := client.Do(request)
	if err != nil {
		return errors.Errorf("[GFWDownloader] http client get %s error: %s", url, err.Error())
	}
	defer response.Body.Close()
	filename := fmt.Sprintf("./resource/gfwrules/%s", path.Base(request.URL.Path))
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Errorf("[GFWDownloader] read %s error: %s", filename, err.Error())
	}
	err = ioutil.WriteFile(filename, body, 0644)
	if err != nil {
		return errors.Errorf("[GFWDownloader] write %s error: %s", filename, err.Error())
	}
	log.Printf("[GFWDownloader] download %s success\n", filename)
	return nil
}
