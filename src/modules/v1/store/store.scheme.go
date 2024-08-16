package store

type findStoreQs struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Order     string `form:"order"`
	Direction string `form:"direction"`
}

type storeBody struct {
	Name    string `binding:"required" json:"name" mod:"trim"`
	Phone   string `binding:"required,numeric" json:"phone" mod:"trim"`
	Address string `binding:"required" json:"address" mod:"trim"`
	WardId  uint   `binding:"required" json:"ward_id"`
	OwnerId uint   `binding:"required" json:"owner_id"`
}
