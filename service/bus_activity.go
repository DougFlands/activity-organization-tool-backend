package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBusActivity
//@description: 创建BusActivity记录
//@param: busAct model.BusActivity
//@return: err error

func CreateBusActivity(busAct model.BusActivity) (err error) {
	err = global.GVA_DB.Create(&busAct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBusActivity
//@description: 删除BusActivity记录
//@param: busAct model.BusActivity
//@return: err error

func DeleteBusActivity(busAct model.BusActivity) (err error) {
	err = global.GVA_DB.Delete(&busAct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBusActivityByIds
//@description: 批量删除BusActivity记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBusActivityByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BusActivity{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBusActivity
//@description: 更新BusActivity记录
//@param: busAct *model.BusActivity
//@return: err error

func UpdateBusActivity(busAct model.BusActivity) (err error) {
	err = global.GVA_DB.Save(&busAct).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivity
//@description: 根据id获取BusActivity记录
//@param: id uint
//@return: err error, busAct model.BusActivity

func GetBusActivity(id uint) (err error, busAct model.BusActivity) {
	err = global.GVA_DB.Where("id = ?", id).First(&busAct).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivityInfoList
//@description: 分页获取BusActivity记录
//@param: info request.BusActivitySearch
//@return: err error, list interface{}, total int64

func GetBusActivityInfoList(info request.BusActivitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusActivity{})
	var busActs []model.BusActivity
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.User)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busActs).Error
	return err, busActs, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivityInfoList
//@description: 分页获取BusActivity记录
//@param: info request.BusActivitySearch
//@return: err error, list interface{}, total int64

func GetBusInvolvedActivityList(info request.BusInvolvedActivitySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})
	var busInvolvedActs []model.BusInvolvedActivitys
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("user_id = ? AND status = 1", info.User)

	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busInvolvedActs).Error

	return err, busInvolvedActs, total
}

func InvolvedOrExitActivities(busAct model.BusInvolvedActivitys) (err error) {
	var searchBusAct model.BusInvolvedActivitys
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if err := db.Where("user_id = ? AND activity_id = ?", busAct.User, busAct.Activity).First(&searchBusAct).Error; err != nil {
		err = global.GVA_DB.Create(&busAct).Error
		return err
	}
	busAct.CreatedAt = searchBusAct.CreatedAt
	busAct.ID = searchBusAct.ID
	utils.ToolJsonFmt(busAct)
	utils.ToolJsonFmt(searchBusAct)
	err = global.GVA_DB.Save(&busAct).Error
	return err

}
