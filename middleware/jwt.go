package middleware

import (
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
)

func RegAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("x-user-id")
		accept := service.FindUser(userId)
		if !accept {
			response.FailWithAuthUserMessage("用户未注册", c)
			c.Abort()
		}
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("x-user-id")
		accept := service.FindAdminUser(userId)
		if !accept {
			response.FailWithAuthAdminMessage("用户非管理员", c)
			c.Abort()
		}
		c.Next()
	}
}

func HighestAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Request.Header.Get("x-user-id")
		accept := service.FindHighestAndAdminUser(userId)
		if !accept {
			response.FailWithAuthAdminMessage("用户非最高管理员", c)
			c.Abort()
		}
		c.Next()
	}
}
