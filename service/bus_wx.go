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
		"wxmsg":      wxmsg,
		"TemplateID": m.TemplateID,
		"url":        "/pages/activity/detail?id=" + wxmsg.ActivityId,
	})

	for _, V := range m.TemplateID {
		msg := &subscribe.Message{
			ToUser:     wxmsg.UserOpenId,
			TemplateID: V,
			Data: map[string]*subscribe.DataItem{
				// 活动名称
				"thing1": {
					Value: wxmsg.ActivityName,
				},
				// 活动时间
				"time2": {
					Value: wxmsg.ActivityTime,
				},
				// 提醒内容
				"thing3": {
					Value: wxmsg.Content,
				},
			},
			Page: "/pages/activity/detail?id=" + wxmsg.ActivityId,
		}
		err := sub.Send(msg)
		fmt.Println(err)
		if err == nil {
			break
		}
	}

}
