package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"time"
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

func GetBusActivity(id uint) (err error, busActDetail model.BusActivityDetail) {
	var busAct model.BusActivity
	// 创建db
	db := global.GVA_DB.Model(&model.BusActivity{})
	// 如果有条件搜索 下方会自动创建搜索语句
	// 查询当前用户创建的活动
	db = db.Where("id = ?", id).Preload("User").Preload("BusGame")
	err = db.Find(&busAct).Error
	// 搜索参与人数
	participants, _, _ := findInvolvedParticipants(0, id)
	busAct.Participants = participants
	userList, err := findInvolvedUser(id)
	busActDetail.BusActivity = busAct
	busActDetail.UserList = userList
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivityInfoList
//@description: 分页获取BusActivity记录
//@param: info request.BusActivitySearch
//@return: err error, list interface{}, total int64

func GetBusActivityInfoList(info request.BusActivitySearch) (err error, list []response.BusInvolvedActivitysRes, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusActivity{})
	var busActs []model.BusActivity
	// 如果有条件搜索 下方会自动创建搜索语句
	// 查询当前用户创建的活动
	if info.UserId != 0 {
		db = db.Where("user_id = ?", info.UserId)
	} else {
		db = db.Where("date_time >= ?", time.Now().Format("2006-01-02 15:04:05"))
	}
	db = db.Order("date_time Desc").Preload("User").Preload("BusGame")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busActs).Error

	// 搜索参与人数
	for i := 0; i < len(busActs); i++ {
		participants, _, _ := findInvolvedParticipants(0, busActs[i].ID)
		busActs[i].Participants = participants
		IsInvolved := findIsInvolved(uint(busActs[i].ID), busActs[i].UserId)
		list = append(list, response.BusInvolvedActivitysRes{
			BusActivity: busActs[i],
			IsInvolved:  IsInvolved,
		})
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivityInfoList
//@description: 分页获取已参与的BusActivity记录
//@param: info request.BusActivitySearch
//@return: err error, list interface{}, total int64

func GetBusInvolvedActivityList(info request.BusInvolvedActivitySearch) (err error, busInvolvedActs []model.BusInvolvedActivitys, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("user_id = ? AND status = 1", info.UserId).Order("updated_at Desc").Preload("User").Preload("Activity").Preload("Activity.BusGame").Preload("Activity.User")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busInvolvedActs).Error
	// 搜索参与人数
	for i := 0; i < len(busInvolvedActs); i++ {
		participants, _, _ := findInvolvedParticipants(0, uint(busInvolvedActs[i].ActivityId))
		busInvolvedActs[i].Activity.Participants = participants
	}
	return err, busInvolvedActs, total
}

func InvolvedOrExitActivities(busAct model.BusInvolvedActivitys) (err error) {
	var searchBusAct model.BusInvolvedActivitys
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})
	// 如果有条件搜索 下方会自动创建搜索语句
	if err := db.Where("user_id = ? AND activity_id = ?", busAct.UserId, busAct.ActivityId).Preload("Activity.BusGame").First(&searchBusAct).Error; err != nil {
		err = global.GVA_DB.Create(&busAct).Error
		return err
	}
	busAct.CreatedAt = searchBusAct.CreatedAt
	busAct.ID = searchBusAct.ID

	if busAct.Status == 1 {
		participants, _, _ := findInvolvedParticipants(busAct.UserId, uint(busAct.ActivityId))

		if participants != 0 {
			return errors.New("已参与过该活动")
		}

		involvedActivityParticipants, _, _ := findInvolvedParticipants(0, uint(busAct.ActivityId))
		if involvedActivityParticipants != 0 && involvedActivityParticipants >= searchBusAct.Activity.BusGame.PeopleNum {
			return errors.New("参与人数达到上限")
		}

	}
	err = global.GVA_DB.Save(&busAct).Error
	return err

}

// 搜索参与人数
func findInvolvedParticipants(userId int, activityId uint) (participant int, involvedActivity model.BusInvolvedActivitys, err error) {
	var participants int64
	involvedDb := global.GVA_DB.Model(&involvedActivity)
	if userId > 0 {
		involvedDb = involvedDb.Where("user_id = ?", userId)
	}
	if activityId > 0 {
		involvedDb = involvedDb.Where("activity_id = ?", activityId)
	}
	involvedDb = involvedDb.Where("status = 1").Preload("User").Preload("Activity")
	err = involvedDb.Count(&participants).Error
	return int(participants), involvedActivity, err
}

// 搜索参与的人
func findInvolvedUser(activityId uint) (involvedUser []model.SysUserInfo, err error) {
	var involvedActivity []model.BusInvolvedActivitys

	involvedActivityDb := global.GVA_DB.Model(&involvedActivity)
	involvedActivityDb = involvedActivityDb.Where("activity_id = ? and status = 1", activityId).Preload("User").Find(&involvedActivity)
	err = involvedActivityDb.Find(&involvedActivity).Error

	for i := 0; i < len(involvedActivity); i++ {
		involvedUser = append(involvedUser, involvedActivity[i].User)
	}
	return
}

// 搜索是否已参与
func findIsInvolved(activityId uint, userId int) (isInvolved bool) {
	var involvedActivity model.BusInvolvedActivitys
	involvedActivityDb := global.GVA_DB.Model(&involvedActivity)
	_ = involvedActivityDb.Where("activity_id = ? and user_id = ? and status = 1", activityId, userId).First(&involvedActivity).Error
	return involvedActivity.UserId != 0
}
