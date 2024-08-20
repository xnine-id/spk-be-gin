package installation

import (
	"mime/multipart"

	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
)

type findInstallationQs struct {
	pagination.QS
}

type installationBody struct {
	SpkNumber        string `binding:"required" json:"spk_number"`
	SpkDate          string `binding:"required" json:"spk_date"`
	StoreId          uint   `binding:"required" json:"store_id"`
	InstallationDate string `binding:"required" json:"installation_date"`
	SalesId          uint   `binding:"required" json:"sales_id"`
	Status           *bool  `binding:"required" json:"status"`
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
