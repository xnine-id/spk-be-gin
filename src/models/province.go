package models

type Province struct {
	PK
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Timestamps

	// Relations
	Regencies *[]Regency `gorm:"foreignKey:province_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"regencies,omitempty"`
}

func (Province) TableName() string {
	return "mst_provinces"
}
