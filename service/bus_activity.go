package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"strconv"
	"time"

	"gorm.io/gorm"
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

	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})
	db = db.Delete(model.BusInvolvedActivitys{}, "activity_id = ?", busAct.ID)

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

func GetBusActivity(id uint, userId int) (err error, busActDetail model.BusActivityDetail) {
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
	busActDetail.UserList = userList
	busActDetail.BusActivity = busAct
	IsInvolved := findIsInvolved(id, userId)
	busActDetail.IsInvolved = IsInvolved
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBusActivityInfoList
//@description: 分页获取BusActivity记录
//@param: info request.BusActivitySearch
//@return: err error, list interface{}, total int64

func GetBusActivityInfoList(info request.BusActivitySearch, userId int) (err error, list []response.BusInvolvedActivitysRes, total int64) {
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
		// 查询所有活动，此时需要过滤
		db = db.Where("date_time >= ?", time.Now().Format("2006-01-02 15:04:05"))
	}
	db = db.Order("date_time asc").Preload("User").Preload("BusGame")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busActs).Error

	// 搜索参与人数
	for i := 0; i < len(busActs); i++ {
		participants, _, _ := findInvolvedParticipants(0, busActs[i].ID)
		busActs[i].Participants = participants

		IsInvolved := findIsInvolved(uint(busActs[i].ID), int(userId))
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
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{}).Unscoped()
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("user_id = ? AND status = 1", info.UserId).Order("updated_at Desc").Preload("User").Preload("Activity", func(db *gorm.DB) *gorm.DB {
		return db.Unscoped()
	}).Preload("Activity.BusGame").Preload("Activity.User")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&busInvolvedActs).Error
	// 搜索参与人数
	for i := 0; i < len(busInvolvedActs); i++ {
		participants, _, _ := findInvolvedParticipants(0, uint(busInvolvedActs[i].ActivityId))
		busInvolvedActs[i].Activity.Participants = participants
	}
	return err, busInvolvedActs, total
}

// 参与或退出活动
func InvolvedOrExitActivities(busAct model.BusInvolvedActivitys) (err error) {
	// 新建一个，搜索时不能覆盖 busAct 的数据
	var searchBusAct model.BusInvolvedActivitys
	db := global.GVA_DB.Model(&model.BusInvolvedActivitys{})

	// 搜索该用户是否已参与
	participants, _, _ := findInvolvedParticipants(busAct.UserId, uint(busAct.ActivityId))
	if participants == 0 {
		// 是否参与标识，创建一个未参与的记录
		isInvolved := false
		if busAct.Status == 1 {
			busAct.Status = 0
			isInvolved = true
		}
		hassInvolved := findHasInvolvedStatus(uint(busAct.ActivityId), busAct.UserId)
		if !hassInvolved {
			if err := global.GVA_DB.Save(&busAct).Error; err != nil {
				return err
			}
		}

		// 恢复参与字段的数据
		if isInvolved {
			busAct.Status = 1
		}
	}

	// 填充参与的数据
	if err := db.Where("user_id = ? AND activity_id = ?", busAct.UserId, busAct.ActivityId).Preload("Activity.User").Preload("Activity.BusGame").First(&searchBusAct).Error; err != nil {
		return err
	}

	busAct.CreatedAt = searchBusAct.CreatedAt
	busAct.ID = searchBusAct.ID

	// 活动参与人数
	involvedActivityParticipants, _, _ := findInvolvedParticipants(0, uint(busAct.ActivityId))

	// 对应的活动
	activitysDb := global.GVA_DB.Model(&model.BusActivity{})

	if busAct.Status == 1 {

		if participants != 0 {
			return errors.New("已参与过该活动")
		}
		if involvedActivityParticipants != 0 && involvedActivityParticipants >= searchBusAct.Activity.BusGame.PeopleNum+3 {
			return errors.New("参与人数达到上限")
		}

		// 参与成功，刷新人数
		activitysDb.Model(&model.BusActivity{}).Where("id = ?", searchBusAct.ActivityId).Update("participants", involvedActivityParticipants+1)

		// 活动人满
		// >= 判断会导致候补参加时也会通知
		if involvedActivityParticipants+1 == searchBusAct.Activity.BusGame.PeopleNum {
			id := strconv.Itoa(int(searchBusAct.ActivityId))
			sendActReadyMsg(model.WxMsg{
				ActivityId:   id,
				ActivityName: searchBusAct.Activity.BusGame.Name,
				ActivityTime: searchBusAct.Activity.DateTime.String(),
				Content:      "活动人数已满",
				UserOpenId:   searchBusAct.Activity.User.OpenID,
			})
		}

	} else {
		// 退出成功，刷新人数
		activitysDb.Model(&model.BusActivity{}).Where("id = ?", searchBusAct.ActivityId).Update("participants", involvedActivityParticipants-1)
		// 备胎转正
		involvedUser, err := findInvolvedUser(uint(searchBusAct.ActivityId))
		if err != nil {
			return err
		}
		// 退出的是否为玩家而不是备胎
		exitPlayerId := 0
		// 第一个备胎id
		firstSpareTireUserId := "0"
		for i, v := range involvedUser {
			if searchBusAct.Activity.BusGame.PeopleNum == i {
				s := strconv.Itoa(int(v.ID))
				firstSpareTireUserId = s
			}
			if busAct.UserId == int(v.ID) {
				exitPlayerId = busAct.UserId
			}
		}
		fmt.Printf("备胎转正 - 游戏ID: %d 退出玩家: %d, 备胎id: %s", busAct.ActivityId, exitPlayerId, firstSpareTireUserId)
		_, firstSpareTireUser := FindUser(firstSpareTireUserId)
		// 有玩家退出且有备胎转正
		if exitPlayerId != 0 && firstSpareTireUserId != "0" {
			id := strconv.Itoa(int(searchBusAct.ActivityId))
			sendActSpareTireInvolved(model.WxMsg{
				ActivityId:   id,
				ActivityName: searchBusAct.Activity.BusGame.Name,
				Content:      "恭喜！您参与的活动备胎转正！",
				UserOpenId:   firstSpareTireUser.OpenID,
			})
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
	err = involvedActivityDb.Order("updated_at asc").Find(&involvedActivity).Error

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

// 搜索是否有参与记录
func findHasInvolvedStatus(activityId uint, userId int) (hasInvolved bool) {
	var involvedActivity model.BusInvolvedActivitys
	involvedActivityDb := global.GVA_DB.Model(&involvedActivity)
	_ = involvedActivityDb.Where("activity_id = ? and user_id = ?", activityId, userId).First(&involvedActivity).Error
	return involvedActivity.ID != 0
}
