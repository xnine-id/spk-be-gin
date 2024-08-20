package migration

import "github.com/amuhajirs/gin-gorm/src/models"

var Tables = []any{
	&models.Province{},
	&models.Regency{},
	&models.Subdistrict{},
	&models.Ward{},
	&models.User{},
	&models.Token{},
	&models.Sales{},
	&models.Store{},
	&models.Installation{},
}