package source

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysUserInfo{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}

		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}
