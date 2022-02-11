// 自动生成模板BusActivity
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type BusActivity struct {
	global.GVA_MODEL
	GameId   int        `json:"gameId" form:"gameId" gorm:"column:game_id;comment:游戏id;"`
	Location string     `json:"location" form:"location" gorm:"column:location;comment:地点;"`
	Price    string     `json:"price" form:"price" gorm:"column:price;comment:价格;"`
	DateTime *LocalTime `json:"dateTime" form:"dateTime" gorm:"column:date_time;comment:时间;"`
}

func (BusActivity) TableName() string {
	return "bus_activity"
}
