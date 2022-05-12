package router

import (
	v1 "gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserAdminRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.GET("userList", v1.FindUserAndAdminUser) // 查找用户列表
		UserRouter.GET("banUserList", v1.FindBanUserList)   // 黑名单列表
		UserRouter.POST("banUser", v1.HandleBanUser)        // 拉黑
	}
}

func InitUserHighestAdminRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.OperationRecord())
	{
		UserRouter.POST("setAdminAuthority", v1.SetAdminAuthority) // 设置用户权限
	}
}
