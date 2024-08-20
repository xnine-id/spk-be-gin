package models

type Store struct {
	PK
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Phone   string `gorm:"type:varchar(20)" json:"phone"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	WardId  uint   `gorm:"type:bigint" json:"ward_id"`
	Owner   string `gorm:"type:varchar(100)" json:"owner"`
	Timestamps

	// Relations
	Ward *Ward `gorm:"foreignKey:ward_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ward,omitempty"`
}

func (Store) TableName() string {
	return "mst_stores"
}
