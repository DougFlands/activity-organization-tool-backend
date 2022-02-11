package router

import (
	v1 "gin-vue-admin/api/v1"

	"github.com/gin-gonic/gin"
)

func InitInitRouter(Router *gin.RouterGroup) {
	InitRouter := Router.Group("init")
	{
		InitRouter.POST("initdb", v1.InitDB)       // 创建Api
		InitRouter.POST("checkdb", v1.CheckDB)     // 创建Api
		InitRouter.POST("initTable", v1.InitTable) // 创建Api

	}
}
