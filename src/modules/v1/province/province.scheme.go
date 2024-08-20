package province

import "github.com/amuhajirs/gin-gorm/src/helpers/pagination"

type findProvinceQs struct {
	pagination.QS
}

type provinceBody struct {
	Name string `binding:"required" json:"name" mod:"trim"`
}
