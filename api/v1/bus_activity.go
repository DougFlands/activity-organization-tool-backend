package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags BusActivity
// @Summary 创建BusActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusActivity true "创建BusActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busAct/createBusActivity [post]
func CreateBusActivity(c *gin.Context) {
	var busAct model.BusActivity
	_ = c.ShouldBindJSON(&busAct)
	userId := c.GetHeader("x-user-id")
	userIdInt, _ := strconv.Atoi(userId)
	busAct.UserId = userIdInt
	if err := service.CreateBusActivity(busAct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BusActivity
// @Summary 删除BusActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusActivity true "删除BusActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busAct/deleteBusActivity [delete]
func DeleteBusActivity(c *gin.Context) {
	var busAct model.BusActivity
	_ = c.ShouldBindJSON(&busAct)
	if err := service.DeleteBusActivity(busAct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BusActivity
// @Summary 批量删除BusActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busAct/deleteBusActivityByIds [delete]
func DeleteBusActivityByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBusActivityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BusActivity
// @Summary 更新BusActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusActivity true "更新BusActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busAct/updateBusActivity [put]
func UpdateBusActivity(c *gin.Context) {
	var busAct model.BusActivity
	_ = c.ShouldBindJSON(&busAct)
	if err := service.UpdateBusActivity(busAct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BusActivity
// @Summary 用id查询BusActivity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusActivity true "用id查询BusActivity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busAct/findBusActivity [get]
func FindBusActivity(c *gin.Context) {
	var busAct model.BusActivity
	_ = c.ShouldBindQuery(&busAct)
	userId := c.GetHeader("x-user-id")
	userIdRes, _ := strconv.Atoi(userId)

	if err, busAct := service.GetBusActivity(busAct.ID, userIdRes); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"busAct": busAct}, c)
	}
}

// @Tags BusActivity
// @Summary 分页获取BusActivity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BusActivitySearch true "分页获取BusActivity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busAct/getBusActivityList [get]
func GetBusActivityList(c *gin.Context) {
	var pageInfo request.BusActivitySearch
	userId := c.GetHeader("x-user-id")
	userIdRes, _ := strconv.Atoi(userId)

	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetBusActivityInfoList(pageInfo, userIdRes); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags BusActivity
// @Summary 分页获取BusActivity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BusActivitySearch true "分页获取BusActivity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busAct/getBusActivityList [get]
func GetBusInvolvedActivityList(c *gin.Context) {
	var pageInfo request.BusInvolvedActivitySearch
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.GetBusInvolvedActivityList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 参加与退出活动
func InvolvedOrExitActivities(c *gin.Context) {
	var busAct model.BusInvolvedActivitys
	_ = c.ShouldBindJSON(&busAct)
	if err := service.InvolvedOrExitActivities(busAct); err != nil {
		global.GVA_LOG.Error("参与失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), c)
	} else {
		if busAct.Status == 1 {
			response.OkWithMessage("参与成功", c)
		} else {
			response.OkWithMessage("退出成功", c)
		}
	}
}
