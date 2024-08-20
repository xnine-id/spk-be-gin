package ward

import "github.com/amuhajirs/gin-gorm/src/helpers/pagination"

type findWardQs struct {
	pagination.QS

	ProvinceId    string `form:"province_id"`
	RegencyId     string `form:"regency_id"`
	SubdistrictId string `form:"subdistrict_id"`
}

type wardBody struct {
	Name          string `binding:"required" json:"name" mod:"trim"`
	SubdistrictId uint   `binding:"required" json:"subdistrict_id"`
}
