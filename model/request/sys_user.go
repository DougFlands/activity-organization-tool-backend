package request

// User register structure
type Register struct {
	Username    string `json:"userName"`
	Password    string `json:"passWord"`
	NickName    string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg   string `json:"headerImg" gorm:"default:'http://www.henrongyi.top/avatar/lufu.jpg'"`
	AuthorityId string `json:"authorityId" gorm:"default:888"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	UserId  int `json:"userId"`  // 角色ID
	IsAdmin int `json:"isAdmin"` // 管理员等级
}

type UserList struct {
	PageInfo
	UserId int `json:"userId" form:"userId"` //
}

type UserBanList struct {
	PageInfo
	PlayerId int    `json:"playerId" form:"playerId"`
	DmId     int    `json:"dmId" form:"dmId"`
	Status   int    `json:"status" form:"status"`
	Content  string `json:"content" form:"content"`
}
