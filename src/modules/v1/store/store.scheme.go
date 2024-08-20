package store

import "github.com/amuhajirs/gin-gorm/src/helpers/pagination"

type findStoreQs struct {
	pagination.QS
}

type storeBody struct {
	Name    string `binding:"required" json:"name" mod:"trim"`
	Phone   string `binding:"required,numeric" json:"phone" mod:"trim"`
	Address string `binding:"required" json:"address" mod:"trim"`
	WardId  uint   `binding:"required" json:"ward_id"`
	Owner   string `binding:"required" json:"owner"`
}
