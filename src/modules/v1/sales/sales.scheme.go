package sales

import (
	"mime/multipart"

	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
)

type findSalesQs struct {
	pagination.QS
}

type createSalesBody struct {
	Name    string                `binding:"required" json:"name" form:"name" mod:"trim"`
	Email   string                `binding:"required,email" json:"email" form:"email" mod:"trim"`
	Phone   string                `binding:"required,numeric" json:"phone" form:"phone" mod:"trim"`
	Address string                `binding:"required" json:"address" form:"address" mod:"trim"`
	WardId  uint                  `binding:"required" json:"ward_id" form:"ward_id"`
	Photo   *multipart.FileHeader `binding:"required" json:"photo" form:"photo"`
}

type updateSalesBody struct {
	Name    string                `binding:"required" json:"name" form:"name" mod:"trim"`
	Email   string                `binding:"required,email" json:"email" form:"email" mod:"trim"`
	Phone   string                `binding:"required,numeric" json:"phone" form:"phone" mod:"trim"`
	Address string                `binding:"required" json:"address" form:"address" mod:"trim"`
	WardId  uint                  `binding:"required" json:"ward_id" form:"ward_id"`
	Photo   *multipart.FileHeader `json:"photo" form:"photo"`
}
