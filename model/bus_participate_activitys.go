package model

import (
	"gin-vue-admin/global"
)

// 参加活动表
// Status: 0: 未参加 1: 参加
type BusInvolvedActivitys struct {
	global.GVA_MODEL
	UserId     int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`
	ActivityId int `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动ID;"`
	Status     int `json:"status" form:"status" gorm:"column:status;comment:状态;"`
}

func (BusInvolvedActivitys) TableName() string {
	return "bus_involved_activity"
}
