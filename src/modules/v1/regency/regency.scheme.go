package regency

type findRegencyQs struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type regencyBody struct {
	Name       string `binding:"required" json:"name" mod:"trim"`
	ProvinceId uint   `binding:"required" json:"province_id"`
}
