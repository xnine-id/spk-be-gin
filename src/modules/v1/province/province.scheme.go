package province

type findProvinceQs struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type provinceBody struct {
	Name string `binding:"required" json:"name" mod:"trim"`
}
