package request

import "gin-vue-admin/model"

type BusGameSearch struct{
    model.BusGame
    PageInfo
}