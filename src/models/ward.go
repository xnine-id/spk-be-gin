package models

type Ward struct {
	PK
	Name          string `gorm:"type:varchar(100);not null" json:"name"`
	SubdistrictId uint   `gorm:"type:bigint" json:"subdistrict_id"`
	Timestamps

	// Relations
	Subdistrict *Subdistrict `gorm:"foreignKey:subdistrict_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subdistrict"`
}
