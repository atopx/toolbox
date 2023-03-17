package scheduler

import (
	"github.com/gin-gonic/gin"
	"superserver/internal/controller"
)

type Controller struct {
	controller.Controller
}

func NewController(ctx *gin.Context) *Controller {
	return &Controller{Controller: controller.Controller{Params: new(Params), Context: ctx}}
}

type Params struct {
	Context struct {
		DeviceCategory string `json:"device_category"`
		DeviceId       string `json:"device_id"`
		InExp          bool   `json:"in_exp"`
		UserAgent      string `json:"user_agent"`
	} `json:"context"`
	Query   string `json:"query"`
	Request struct {
		Intent struct {
			AppId          string `json:"app_id"`
			Complete       bool   `json:"complete"`
			Confidence     int    `json:"confidence"`
			Domain         string `json:"domain"`
			IsDirectWakeup bool   `json:"is_direct_wakeup"`
			IsSkipQuery    bool   `json:"is_skip_query"`
			NeedFetchToken bool   `json:"need_fetch_token"`
			Query          string `json:"query"`
			RequestType    string `json:"request_type"`
			Score          int    `json:"score"`
			SkillType      string `json:"skillType"`
			SkipQueryModel string `json:"skip_query_model"`
			SkipQuerySeed  string `json:"skip_query_seed"`
			Slots          string `json:"slots"`
			SubDomain      string `json:"sub_domain"`
		} `json:"intent"`
		IsMonitor bool      `json:"is_monitor"`
		Locale    string    `json:"locale"`
		RequestId string    `json:"request_id"`
		SlotInfo  Operation `json:"slot_info"`
		Timestamp int64     `json:"timestamp"`
		Type      int       `json:"type"`
	} `json:"request"`
	Session struct {
		Application struct {
			AppId string `json:"app_id"`
		} `json:"application"`
		Attributes struct {
		} `json:"attributes"`
		IsNew     bool   `json:"is_new"`
		SessionId string `json:"session_id"`
		User      struct {
			Gender      string `json:"gender"`
			IsUserLogin bool   `json:"is_user_login"`
			UserId      string `json:"user_id"`
		} `json:"user"`
	} `json:"session"`
	Version string `json:"version"`
}

type Operation struct {
	IntentName string          `json:"intent_name"`
	Slots      []OperationSlot `json:"slots"`
}

type OperationSlot struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	RawValue string `json:"raw_value"`
}

type Reply struct {
	IsSessionEnd bool           `json:"is_session_end"`
	Version      string         `json:"version"`
	Response     XiaoaiResponse `json:"response"`
}

type XiaoaiResponse struct {
	Confidence    float64     `json:"confidence"`
	OpenMic       bool        `json:"open_mic"`
	ToSpeak       Interactive `json:"to_speak"`
	ToDisplay     Interactive `json:"to_display"`
	NotUnderstand bool        `json:"not_understand"`
}

type Interactive struct {
	Type int    `json:"type"`
	Text string `json:"text"`
}
