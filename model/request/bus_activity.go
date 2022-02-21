package request

import "gin-vue-admin/model"

type BusActivitySearch struct {
	model.BusActivity
	PageInfo
}

type BusInvolvedActivitySearch struct {
	model.BusInvolvedActivitys
	PageInfo
	UserId int `json:"userId" form:"userId"`
}
