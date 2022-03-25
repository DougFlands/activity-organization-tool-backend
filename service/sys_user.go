package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"

	uuid "github.com/satori/go.uuid"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Register
//@description: 用户注册
//@param: u model.SysUser
//@return: err error, userInter model.SysUser

func Login(u model.SysLoginInfo) (sysUserInfo model.SysUserInfo, err error) {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	m := global.GVA_CONFIG.Wx
	cfg := &miniConfig.Config{
		AppID:     m.AppID,
		AppSecret: m.AppSecret,
		Cache:     memory,
	}
	mini := wc.GetMiniProgram(cfg)
	auth := mini.GetAuth()
	session, err := auth.Code2Session(u.Code)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
	userInfo := model.SysUserInfo{
		OpenID:    session.OpenID,
		AvatarUrl: u.AvatarUrl,
		NickName:  u.NickName,
	}

	utils.ToolJsonFmt(userInfo)

	// 搜索是否存在数据
	if !errors.Is(global.GVA_DB.Where("openID = ?", userInfo.OpenID).First(&userInfo).Error, gorm.ErrRecordNotFound) {
		// 存在则更新
		err = global.GVA_DB.Where("openID = ?", userInfo.OpenID).Updates(&userInfo).Error
		return userInfo, err
	}

	// 不存在则注册
	err = global.GVA_DB.Create(&userInfo).Error
	return userInfo, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetUserAuthority
//@description: 设置一个用户的权限
//@param: uuid uuid.UUID, authorityId string
//@return: err error

func SetUserAuthority(uuid uuid.UUID, authorityId string) (err error) {
	// err = global.GVA_DB.Where("uuid = ?", uuid).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

func FindUser(userId string) (accept bool) {
	var user model.SysUserInfo
	if err := global.GVA_DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}

func FindAdminUser(userId string) (accept bool) {
	var user model.SysUserInfo
	if err := global.GVA_DB.Where("id = ? and isAdmin = 1", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}
