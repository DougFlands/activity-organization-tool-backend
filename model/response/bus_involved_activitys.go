package response

import "gin-vue-admin/model"

// 参加活动
type BusInvolvedActivitysRes struct {
	model.BusActivity
	IsInvolved bool `json:"isInvolved"`
}
