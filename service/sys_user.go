package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"

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
		utils.ToolJsonFmt(m)
		fmt.Printf("err: %v", err)
		return model.SysUserInfo{}, err
	}
	userInfo := model.SysUserInfo{
		OpenID: session.OpenID,
	}

	// 搜索是否存在数据 TODO: 锁表
	if !errors.Is(global.GVA_DB.Where("openID = ?", userInfo.OpenID).First(&userInfo).Error, gorm.ErrRecordNotFound) {
		// 存在则更新
		userInfo.AvatarUrl = u.AvatarUrl
		userInfo.NickName = u.NickName
		err = global.GVA_DB.Where("openID = ?", userInfo.OpenID).Updates(userInfo).Error
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

func SetUserAuthority(id int, isAdmin int) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.SysUserInfo{}).Update("isAdmin", isAdmin).Error
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
	if err := global.GVA_DB.Where("id = ? and (isAdmin = 1 OR isAdmin = 2)", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}

// 找最高级管理员
func FindHighestAndAdminUser(userId string) (accept bool) {
	var user model.SysUserInfo
	if err := global.GVA_DB.Where("id = ? and isAdmin = 2", userId).First(&user).Error; err != nil {
		return false
	}
	return true
}

// 找用户和1级管理员
func FindUserAndAdminUser(info request.UserList) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.SysUserInfo{})
	var busActs []model.SysUserInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("Not isAdmin = 2")
	if info.UserId != 0 {
		db = db.Where("id = ?", info.UserId)
	}
	db = db.Order("created_at Desc")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busActs).Error
	return err, busActs, total
}
