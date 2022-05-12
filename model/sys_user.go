package model

import (
	"gin-vue-admin/global"
)

type SysLoginInfo struct {
	Code      string `json:"code" form:"code" gorm:"column:code;comment:code;"`
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl" gorm:"column:avatarUrl;comment:用户头像图片的 URL;"`
	NickName  string `json:"nickName" form:"nickName" gorm:"column:nickName;comment:用户昵称;"`
}

type SysUserInfo struct {
	global.GVA_MODEL
	OpenID    string `json:"openID" form:"openID" gorm:"column:openID;comment:openID;"`
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl" gorm:"column:avatarUrl;comment:用户头像图片的 URL;"`
	NickName  string `json:"nickName" form:"nickName" gorm:"column:nickName;comment:用户昵称;"`
	IsAdmin   int    `json:"isAdmin" form:"isAdmin" gorm:"column:isAdmin;comment:管理员;"`
}

type SysBanUserInfo struct {
	global.GVA_MODEL

	PlayerId int         `json:"playerId" form:"playerId" gorm:"column:playerId;comment:playerId;"`
	Player   SysUserInfo `json:"player"  gorm:"foreignKey:ID;references:PlayerId;comment:玩家ID"`

	DmId int         `json:"dmId" form:"dmId" gorm:"column:dmId;comment:dmId;"`
	Dm   SysUserInfo `json:"dm"  gorm:"foreignKey:ID;references:DmId;comment:DM"`

	Status  int    `json:"status" form:"status" gorm:"column:status;comment:状态 0: 取消拉黑 1: 拉黑 2: 全局拉黑;"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:拉黑原因;"`
}

func (SysBanUserInfo) TableName() string {
	return "sys_ban_user_info"
}
