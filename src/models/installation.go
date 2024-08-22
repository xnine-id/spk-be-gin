package models

import "time"

type Installation struct {
	PK
	SpkNumber        string     `gorm:"type:varchar(50);unique;not null" json:"spk_number"`
	SpkDate          *time.Time `gorm:"type:date" json:"spk_date"`
	StoreId          uint       `gorm:"type:bigint" json:"store_id"`
	InstallationDate *time.Time `gorm:"type:date" json:"installation_date"`
	SalesId          uint       `gorm:"type:bigint" json:"sales_id"`
	Status           bool       `gorm:"type:boolean;default:false;not null" json:"status"`
	Timestamps

	// Relations
	Store  *Store               `gorm:"foreignKey:store_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"store,omitempty"`
	Sales  *Sales               `gorm:"foreignKey:sales_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sales,omitempty"`
	Images *[]InstallationImage `gorm:"foreignKey:installation_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"images,omitempty"`
}

func (Installation) TableName() string {
	return "trx_installations"
}
