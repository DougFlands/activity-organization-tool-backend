package v1

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags BusGame
// @Summary 创建BusGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGame true "创建BusGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGame/createBusGame [post]
func CreateBusGame(c *gin.Context) {
	var busGame model.BusGame
	_ = c.ShouldBindJSON(&busGame)
	utils.ToolJsonFmt(busGame)

	if err := service.CreateBusGame(busGame); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BusGame
// @Summary 删除BusGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGame true "删除BusGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /busGame/deleteBusGame [delete]
func DeleteBusGame(c *gin.Context) {
	var busGame model.BusGame
	_ = c.ShouldBindJSON(&busGame)
	if err := service.DeleteBusGame(busGame); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BusGame
// @Summary 批量删除BusGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BusGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /busGame/deleteBusGameByIds [delete]
func DeleteBusGameByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBusGameByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BusGame
// @Summary 更新BusGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGame true "更新BusGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /busGame/updateBusGame [put]
func UpdateBusGame(c *gin.Context) {
	var busGame model.BusGame
	_ = c.ShouldBindJSON(&busGame)
	if err := service.UpdateBusGame(busGame); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BusGame
// @Summary 用id查询BusGame
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BusGame true "用id查询BusGame"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /busGame/findBusGame [get]
func FindBusGame(c *gin.Context) {
	var busGame model.BusGame
	_ = c.ShouldBindQuery(&busGame)
	if err, gameInfo := service.GetBusGame(busGame.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"gameInfo": gameInfo}, c)
	}
}

// @Tags BusGame
// @Summary 分页获取BusGame列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BusGameSearch true "分页获取BusGame列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /busGame/getBusGameList [get]
func GetBusGameList(c *gin.Context) {
	var pageInfo request.BusGameSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBusGameInfoList(pageInfo); err != nil {
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
