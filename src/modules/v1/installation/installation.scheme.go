package installation

import (
	"mime/multipart"

	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
)

type findInstallationQs struct {
	pagination.QS
}

type installationBody struct {
	SpkNumber        string                  `binding:"required" json:"spk_number" form:"spk_number"`
	SpkDate          string                  `binding:"required" json:"spk_date" form:"spk_date"`
	StoreId          uint                    `binding:"required" json:"store_id" form:"store_id"`
	InstallationDate string                  `binding:"required" json:"installation_date" form:"installation_date"`
	SalesId          uint                    `binding:"required" json:"sales_id" form:"sales_id"`
	Status           *bool                   `binding:"required" json:"status" form:"status"`
	Images           []*multipart.FileHeader `binding:"required,len=4" json:"images" form:"images"`
}

type updateInstallationBody struct {
	SpkNumber        string                  `binding:"required" json:"spk_number" form:"spk_number"`
	SpkDate          string                  `binding:"required" json:"spk_date" form:"spk_date"`
	StoreId          uint                    `binding:"required" json:"store_id" form:"store_id"`
	InstallationDate string                  `binding:"required" json:"installation_date" form:"installation_date"`
	SalesId          uint                    `binding:"required" json:"sales_id" form:"sales_id"`
	Status           *bool                   `binding:"required" json:"status" form:"status"`
	Images           []*multipart.FileHeader `json:"images" form:"images"`
	ImageIds         []uint                  `json:"image_ids" form:"image_ids"`
}

type importInstallationBody struct {
	File *multipart.FileHeader `binding:"required" json:"file" form:"file"`
	ColumnMap
}

type ColumnMap struct {
	SpkNumber string `binding:"required" json:"spk_number" form:"spk_number"`
	SpkDate   string `binding:"required" json:"spk_date" form:"spk_date"`
	StoreId   string `binding:"required" json:"store_id" form:"store_id"`
}
