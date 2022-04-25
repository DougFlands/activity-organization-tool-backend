package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	subscribe "github.com/silenceper/wechat/v2/miniprogram/subscribe"
)

func creatWxmsg() (sub *subscribe.Subscribe) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	m := global.GVA_CONFIG.Wx
	cfg := &miniConfig.Config{
		AppID:     m.AppID,
		AppSecret: m.AppSecret,
		Cache:     memory,
	}

	mini := wc.GetMiniProgram(cfg)
	sub = mini.GetSubscribe()
	return sub
}

func sendActSpareTireInvolved(wxmsg model.WxMsg) {
	sub := creatWxmsg()
	// 日志
	fmt.Println("备胎转正")
	utils.ToolJsonFmt(map[string]interface{}{
		"wxmsg":                        wxmsg,
		"TemplateActSpareTireInvolved": global.GVA_CONFIG.Wx.TemplateActSpareTireInvolved,
		"url":                          "/pages/activity/detail?id=" + wxmsg.ActivityId,
	})
	for _, V := range global.GVA_CONFIG.Wx.TemplateActSpareTireInvolved {
		msg := &subscribe.Message{
			ToUser:     wxmsg.UserOpenId,
			TemplateID: V.Id,
			Data: map[string]*subscribe.DataItem{
				// 活动名称
				V.ActivityName: {
					Value: wxmsg.ActivityName,
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

func sendActReadyMsg(wxmsg model.WxMsg) {
	sub := creatWxmsg()
	// 日志
	fmt.Println("活动人齐")
	utils.ToolJsonFmt(map[string]interface{}{
		"wxmsg":            wxmsg,
		"TemplateActReady": global.GVA_CONFIG.Wx.TemplateActReady,
		"url":              "/pages/activity/detail?id=" + wxmsg.ActivityId,
	})

	for _, V := range global.GVA_CONFIG.Wx.TemplateActReady {
		msg := &subscribe.Message{
			ToUser:     wxmsg.UserOpenId,
			TemplateID: V.Id,
			Data: map[string]*subscribe.DataItem{
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
