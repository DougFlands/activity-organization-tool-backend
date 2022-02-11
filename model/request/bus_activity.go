package request

import "gin-vue-admin/model"

type BusActivitySearch struct {
	model.BusActivity
	PageInfo
}
