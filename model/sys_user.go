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
