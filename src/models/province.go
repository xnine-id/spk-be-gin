package models

type Province struct {
	PK
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Timestamps
}
