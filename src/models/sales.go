package models

type Sales struct {
	PK
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Email   string `gorm:"type:varchar(50)" json:"email"`
	Phone   string `gorm:"type:varchar(20)" json:"phone"`
	Address string `gorm:"type:varchar(255)" json:"address"`
	WardId  uint   `gorm:"type:bigint" json:"ward_id"`
	Photo   string `gorm:"type:varchar(255)" json:"photo"`
	Timestamps

	// Relations
	Ward *Ward `gorm:"foreignKey:ward_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"ward,omitempty"`
}
