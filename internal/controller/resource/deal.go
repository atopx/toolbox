package resource

import (
	"io/ioutil"
	"net/http"
)

func (ctl *Controller) Deal() {
	var filepath string
	switch ctl.param.Filename {
	case "direct":
		filepath = "./resource/gfwrules/direct.txt"
	case "proxy":
		filepath = "./resource/gfwrules/proxy.txt"
	case "reject":
		filepath = "./resource/gfwrules/reject.txt"
	case "private":
		filepath = "./resource/gfwrules/private.txt"
	case "apple":
		filepath = "./resource/gfwrules/apple.txt"
	case "icloud":
		filepath = "./resource/gfwrules/icloud.txt"
	case "telegramcidr":
		filepath = "./resource/gfwrules/telegramcidr.txt"
	case "lancidr":
		filepath = "./resource/gfwrules/lancidr.txt"
	case "cncidr":
		filepath = "./resource/gfwrules/cncidr.txt"
	case "applications":
		filepath = "./resource/gfwrules/applications.txt"
	}
	data, _ := ioutil.ReadFile(filepath)
	ctl.Context.String(http.StatusOK, string(data))
}
