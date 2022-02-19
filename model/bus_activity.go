package model

import (
	"gin-vue-admin/global"
)

type BusActivity struct {
	global.GVA_MODEL

	BusGame BusGame `json:"busGame"  gorm:"foreignKey:ID;references:GameId;comment:游戏"`
	GameId  int     `json:"gameId" form:"gameId" gorm:"column:game_id;comment:游戏id;"`

	User   SysUserInfo `json:"user"  gorm:"foreignKey:ID;references:UserId;comment:用户角色"`
	UserId int         `json:"userId" gorm:"comment:用户id"`

	Location     string     `json:"location" form:"location" gorm:"column:location;comment:地点;"`
	Price        string     `json:"price" form:"price" gorm:"column:price;comment:价格;"`
	Participants int        `json:"participants" form:"participants" gorm:"column:participants;comment:参加人数;default 0"`
	DateTime     *LocalTime `json:"dateTime" form:"dateTime" gorm:"column:date_time;comment:时间;"`
}

func (BusActivity) TableName() string {
	return "bus_activity"
}
