package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("")
	{
		router.InitBaseRouter(PublicGroup)              // 注册基础功能路由 不做鉴权
		router.InitBusActivityPublicRouter(PublicGroup) // 注册活动路由
		router.InitBusGamePublicRouter(PublicGroup)     // 注册游戏路由
		router.InitInitRouter(PublicGroup)              // 自动初始化相关
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.RegAuth())
	{
		router.InitUserRouter(PrivateGroup) // 注册用户路由
	}

	AdminGroup := Router.Group("")
	AdminGroup.Use(middleware.AdminAuth())
	{
		router.InitBusGameAdminRouter(AdminGroup)      // 游戏管理
		router.InitBusActivityAdminRouter(PublicGroup) // 注册活动路由
	}
	global.GVA_LOG.Info("router register success")
	return Router
}
