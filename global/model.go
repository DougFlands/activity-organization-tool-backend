package global

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `json:"id" form:"id" gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      `json:"createdTime" form:"createdTime"` // 创建时间
	UpdatedAt time.Time      `json:"updatedTime" form:"updatedTime"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                 // 删除时间
}
