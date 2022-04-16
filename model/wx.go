// 自动生成模板BusGame
package model

// 如果含有time.Time 请自行import time包
type WxMsg struct {
	ActivityId   string `json:"activityId"`   // 活动名称
	ActivityName string `json:"activityName"` // 活动名称
	ActivityTime string `json:"activityTime"` // 活动时间
	Content      string `json:"content"`      // 消息内容
	UserOpenId   string `json:"userOpenId"`   // 微信openid
}
