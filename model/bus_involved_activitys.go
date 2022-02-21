package model

import (
	"gin-vue-admin/global"
)

// 参加活动表
// Status: 0: 未参加 1: 参加
type BusInvolvedActivitys struct {
	global.GVA_MODEL

	User   SysUserInfo `json:"user" gorm:"foreignKey:ID ;references:UserId;comment:用户角色"`
	UserId int         `json:"userId" gorm:"comment:用户id"`

	Activity   BusActivity `json:"busActivity" gorm:"foreignKey:ID;references:ActivityId;comment:用户角色"`
	ActivityId int         `json:"activityId" gorm:"comment:活动ID"`

	Status int `json:"status" form:"status" gorm:"column:status;comment:状态;"`
}

func (BusInvolvedActivitys) TableName() string {
	return "bus_involved_activity"
}
