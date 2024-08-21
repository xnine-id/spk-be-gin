package models

import (
	"time"
)

type PK struct {
	Id *uint `gorm:"type:bigint;primaryKey" json:"id"`
}

type Timestamps struct {
	UpdatedAt time.Time      `json:"updated_at"`
	CreatedAt time.Time      `json:"created_at"`
}
