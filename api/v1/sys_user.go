package v1

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var wxLogin model.SysLoginInfo
	_ = c.ShouldBindJSON(&wxLogin)
	sysUserInfo, err := service.Login(wxLogin)
	if err != nil {
		fmt.Printf("err: %v", err)
		response.FailWithMessage("登陆失败", c)
		return
	}
	response.OkWithDetailed(response.SysUserResponse{
		User: sysUserInfo,
	}, "获取成功", c)
}

// 查找用户和1级管理员
func FindUserAndAdminUser(c *gin.Context) {
	var pageInfo request.UserList
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.FindUserAndAdminUser(pageInfo); err != nil {
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

// @Tags SysUser
// @Summary 设置用户权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SetUserAuth true "用户UUID, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /user/setUserAuthority [post]
func SetAdminAuthority(c *gin.Context) {
	var sua request.SetUserAuth
	_ = c.ShouldBindJSON(&sua)
	// if UserVerifyErr := utils.Verify(sua, utils.SetUserAuthorityVerify); UserVerifyErr != nil {
	// 	response.FailWithMessage(UserVerifyErr.Error(), c)
	// 	return
	// }
	if err := service.SetUserAuthority(sua.UserId, sua.IsAdmin); err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func FindBanUserList(c *gin.Context) {
	var pageInfo request.UserBanList
	_ = c.ShouldBindQuery(&pageInfo)

	if err, list, total := service.FindBanUserList(pageInfo); err != nil {
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

func HandleBanUser(c *gin.Context) {
	var params model.SysBanUserInfo
	_ = c.ShouldBindJSON(&params)
	if err := service.HandleBanUser(params); err != nil {
		global.GVA_LOG.Error("操作失败!"+err.Error(), zap.Any("err", err))
		response.FailWithMessage("操作失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("操作成功", c)
	}
}
