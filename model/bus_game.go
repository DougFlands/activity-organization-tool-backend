// 自动生成模板BusGame
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type BusGame struct {
	global.GVA_MODEL
	Type         int    `json:"type" form:"type" gorm:"column:type;comment:游戏类型 1 剧本 2 桌游;"`
	Name         string `json:"name" form:"name" gorm:"column:name;comment:地点;"`
	Introduction string `json:"introduction" form:"introduction" gorm:"column:introduction;comment:价格;"`
	PeopleNum    int    `json:"peopleNum" form:"peopleNum" gorm:"column:people_num;comment:最大人数;"`
}

func (BusGame) TableName() string {
	return "bus_game"
}
