package subdistrict

type findSubdistrictQs struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type subdistrictBody struct {
	Name      string `binding:"required" json:"name" mod:"trim"`
	RegencyId uint   `binding:"required" json:"regency_id"`
}
