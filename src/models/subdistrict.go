package models

type Subdistrict struct {
	PK
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	RegencyId  uint   `gorm:"type:bigint" json:"regency_id"`
	Timestamps

	// Relations
	Regency  *Regency  `gorm:"foreignKey:regency_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"regency"`
}
