package models

type User struct {
	PK
	Name     string  `gorm:"type:varchar(50);not null" json:"name"`
	Email    string  `gorm:"type:varchar(50);not null;unique" json:"email"`
	Phone    *string `gorm:"type:varchar(20);unique" json:"phone"`
	Password string  `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Avatar   string  `gorm:"type:varchar(255)" json:"avatar,omitempty"`
	Timestamps
}
