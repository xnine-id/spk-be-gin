package models

type Regency struct {
	PK
	Name       string `gorm:"type:varchar(100);not null" json:"name"`
	ProvinceId uint   `gorm:"type:bigint" json:"province_id"`
	Timestamps

	// Relations
	Province     *Province      `gorm:"foreignKey:province_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"province,omitempty"`
	Subdistricts *[]Subdistrict `gorm:"foreignKey:regency_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"subdistricts,omitempty"`
}

func (Regency) TableName() string {
	return "mst_regencies"
}
