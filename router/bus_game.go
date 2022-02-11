package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBusGamePublicRouter(Router *gin.RouterGroup) {
	BusGameRouter := Router.Group("busGame").Use(middleware.OperationRecord())
	{
		BusGameRouter.GET("findBusGame", v1.FindBusGame)       // 根据ID获取BusGame
		BusGameRouter.GET("getBusGameList", v1.GetBusGameList) // 获取BusGame列表
	}
}

func InitBusGameAdminRouter(Router *gin.RouterGroup) {
	BusGameRouter := Router.Group("busGame").Use(middleware.OperationRecord())
	{
		BusGameRouter.POST("createBusGame", v1.CreateBusGame)           // 新建BusGame
		BusGameRouter.POST("deleteBusGame", v1.DeleteBusGame)           // 删除BusGame
		BusGameRouter.POST("deleteBusGameByIds", v1.DeleteBusGameByIds) // 批量删除BusGame
		BusGameRouter.POST("updateBusGame", v1.UpdateBusGame)           // 更新BusGame
	}
}
