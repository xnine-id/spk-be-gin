package ward

type findWardQs struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type wardBody struct {
	Name          string `binding:"required" json:"name" mod:"trim"`
	SubdistrictId uint   `binding:"required" json:"subdistrict_id"`
}
