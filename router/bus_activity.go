package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBusActivityPublicRouter(Router *gin.RouterGroup) {
	BusActivityRouter := Router.Group("busAct").Use(middleware.OperationRecord())
	{
		BusActivityRouter.GET("findBusActivity", v1.FindBusActivity)       // 根据ID获取BusActivity
		BusActivityRouter.GET("getBusActivityList", v1.GetBusActivityList) // 获取BusActivity列表

		BusActivityRouter.POST("involvedOrExitActivities", v1.InvolvedOrExitActivities) // 参加活动
	}
}

func InitBusActivityAdminRouter(Router *gin.RouterGroup) {
	BusActivityRouter := Router.Group("busAct").Use(middleware.OperationRecord())
	{
		BusActivityRouter.POST("createBusActivity", v1.CreateBusActivity)           // 新建BusActivity
		BusActivityRouter.POST("deleteBusActivity", v1.DeleteBusActivity)           // 删除BusActivity
		BusActivityRouter.POST("deleteBusActivityByIds", v1.DeleteBusActivityByIds) // 批量删除BusActivity
		BusActivityRouter.POST("updateBusActivity", v1.UpdateBusActivity)           // 更新BusActivity
	}
}
