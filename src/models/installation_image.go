package models

type InstallationImage struct {
	PK
	InstallationId uint   `gorm:"bigint" json:"installation_id"`
	Image          string `gorm:"type:varchar(255)" json:"image"`
	Timestamps

	// Relations
	Installation *Installation `gorm:"foreignKey:installation_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"installation,omitempty"`
}

func (InstallationImage) TableName() string {
	return "trx_installation_images"
}