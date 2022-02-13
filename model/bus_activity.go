package model

import (
	"gin-vue-admin/global"
)

type BusActivity struct {
	global.GVA_MODEL
	GameId       int        `json:"gameId" form:"gameId" gorm:"column:game_id;comment:游戏id;"`
	UserId       int        `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`
	Location     string     `json:"location" form:"location" gorm:"column:location;comment:地点;"`
	Price        string     `json:"price" form:"price" gorm:"column:price;comment:价格;"`
	Participants int        `json:"participants" form:"participants" gorm:"column:participants;comment:参加人数;default 0"`
	DateTime     *LocalTime `json:"dateTime" form:"dateTime" gorm:"column:date_time;comment:时间;"`
}

func (BusActivity) TableName() string {
	return "bus_activity"
}
