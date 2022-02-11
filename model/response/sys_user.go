package response

import (
	"gin-vue-admin/model"
)

type SysUserResponse struct {
	User model.SysUserInfo `json:"user"`
}
