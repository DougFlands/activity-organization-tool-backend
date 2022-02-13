package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBusGame
//@description: 创建BusGame记录
//@param: busGame model.BusGame
//@return: err error

func CreateBusGame(busGame model.BusGame) (err error) {
	err = global.GVA_DB.Create(&busGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBusGame
//@description: 删除BusGame记录
//@param: busGame model.BusGame
//@return: err error

func DeleteBusGame(busGame model.BusGame) (err error) {
	err = global.GVA_DB.Delete(&busGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBusGameByIds
//@description: 批量删除BusGame记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBusGameByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BusGame{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBusGame
//@description: 更新BusGame记录
//@param: busGame *model.BusGame
//@return: err error

func UpdateBusGame(busGame model.BusGame) (err error) {
	err = global.GVA_DB.Save(&busGame).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusGame
//@description: 根据id获取BusGame记录
//@param: id uint
//@return: err error, busGame model.BusGame

func GetBusGame(id uint) (err error, busGame model.BusGame) {
	err = global.GVA_DB.Where("id = ?", id).First(&busGame).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusGameInfoList
//@description: 分页获取BusGame记录
//@param: info request.BusGameSearch
//@return: err error, list interface{}, total int64

func GetBusGameInfoList(info request.BusGameSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusGame{})
	var busGames []model.BusGame
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Type != 0 {
		db = db.Where("type = ?", info.Type)
	}
	_ = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busGames).Error
	return busGames, total, err
}
