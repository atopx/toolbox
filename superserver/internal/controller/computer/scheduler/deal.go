package scheduler

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"superserver/common/interface/computer_iface"
	"superserver/common/logger"
	"superserver/common/utils"
	"superserver/internal/model/computer"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	device := viper.GetString("xiaoai.device_id")
	userId := viper.GetString("xiaoai.user_id")
	// 如果不是我的设备，那就听歌吧
	if params.Context.DeviceId != device && params.Session.User.UserId != userId {
		ctl.Context.JSON(http.StatusOK, &Reply{
			IsSessionEnd: true,
			Response: XiaoaiResponse{
				Confidence:    1,
				NotUnderstand: false,
				OpenMic:       true,
				ToSpeak:       Interactive{Text: "好的，请欣赏音乐吧"},
			},
		})
		return
	}

	var reply *Reply
	switch params.Request.SlotInfo.IntentName {
	case "computer":
		reply = ctl.scheduler(params.Request.SlotInfo.Slots)
	default:
		reply = &Reply{
			IsSessionEnd: true,
			Response: XiaoaiResponse{
				Confidence:    1,
				NotUnderstand: true,

				OpenMic: false,
			},
		}
	}
	ctl.Context.JSON(http.StatusOK, reply)
}

func (ctl *Controller) scheduler(slots []OperationSlot) *Reply {
	reply := &Reply{
		IsSessionEnd: true,
		Response: XiaoaiResponse{
			Confidence:    1,
			NotUnderstand: false,
			OpenMic:       false,
		},
	}
	var slot OperationSlot
	if len(slots) > 0 {
		slot = slots[0]
	}
	var po computer.Computer
	dao := computer.NewDao(ctl.GetDatabase())
	dao.Connection().Where("name='homepc'").First(&po)

	switch slot.Name {
	case "computer_operation":
		switch slot.Value {
		case "打开":
			if err := utils.ComputerOn(po.Address); err != nil {
				logger.Error(ctl.Context, "打开电脑失败", zap.Error(err))
				reply.Response.ToSpeak.Text = fmt.Sprintf("Sorry, 操作失败了, %s", err.Error())
			} else {
				reply.Response.ToSpeak.Text = "好的，已为您打开电脑"
			}

		case "关闭":
			if err := utils.ComputerOff(po.Username, po.Password, po.LanHostname, 22); err != nil {
				logger.Error(ctl.Context, "关闭电脑失败", zap.Error(err))
				reply.Response.ToSpeak.Text = fmt.Sprintf("Sorry, 操作失败了, %s", err.Error())
			} else {
				reply.Response.ToSpeak.Text = "好的, 已为您关闭电脑"
			}

		case "检测":
			switch utils.ComputerCheck(po.LanHostname, 3389) {
			case computer_iface.ComputerPowerStatus_COMPUTER_POWER_ON:
				reply.Response.ToSpeak.Text = "您的电脑当前是开机状态"
			case computer_iface.ComputerPowerStatus_COMPUTER_POWER_OFF:
				reply.Response.ToSpeak.Text = "您的电脑当前是关机状态"
			default:
				reply.Response.ToSpeak.Text = "Sorry, 无法检测您的电脑状态"
			}
		default:
			reply.Response.NotUnderstand = true
		}
	}
	return reply
}
