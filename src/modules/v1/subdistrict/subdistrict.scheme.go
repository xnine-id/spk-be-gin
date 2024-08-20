package subdistrict

import "github.com/amuhajirs/gin-gorm/src/helpers/pagination"

type findSubdistrictQs struct {
	pagination.QS
	ProvinceId string `form:"province_id"`
	RegencyId  string `form:"regency_id"`
}

type subdistrictBody struct {
	Name      string `binding:"required" json:"name" mod:"trim"`
	RegencyId uint   `binding:"required" json:"regency_id"`
}
