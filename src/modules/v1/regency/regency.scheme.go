package regency

import "github.com/amuhajirs/gin-gorm/src/helpers/pagination"

type findRegencyQs struct {
	pagination.QS
	ProvinceId string `form:"province_id"`
}

type regencyBody struct {
	Name       string `binding:"required" json:"name" mod:"trim"`
	ProvinceId uint   `binding:"required" json:"province_id"`
}
