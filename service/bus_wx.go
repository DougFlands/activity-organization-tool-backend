package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	subscribe "github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

func sendMsg(wxmsg model.WxMsg) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	m := global.GVA_CONFIG.Wx
	cfg := &miniConfig.Config{
		AppID:     m.AppID,
		AppSecret: m.AppSecret,
		Cache:     memory,
	}

	mini := wc.GetMiniProgram(cfg)
	sub := mini.GetSubscribe()

	// 日志
	utils.ToolJsonFmt(map[string]interface{}{
		"wxmsg":       wxmsg,
		"TemplateMsg": m.TemplateMsg,
		"url":         "/pages/activity/detail?id=" + wxmsg.ActivityId,
	})

	for _, V := range m.TemplateMsg {
		msg := &subscribe.Message{
			ToUser:     wxmsg.UserOpenId,
			TemplateID: V.Id,
			Data: map[string]*subscribe.DataItem{
				// 模板1
				// 活动名称
				V.ActivityName: {
					Value: wxmsg.ActivityName,
				},
				// 活动时间
				V.ActivityTime: {
					Value: wxmsg.ActivityTime,
				},
				// 提醒内容
				V.ActivityContent: {
					Value: wxmsg.Content,
				},
			},
			Page: "/pages/activity/detail?id=" + wxmsg.ActivityId,
		}
		err := sub.Send(msg)
		if err == nil {
			break
		}
	}

}
