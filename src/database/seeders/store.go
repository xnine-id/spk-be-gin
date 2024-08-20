package seeders

import (
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/brianvoe/gofakeit/v7"
)

func StoreSeeder() {
	var stores []*models.Store
	var wardIds []uint

	database.DB.Model(models.Ward{}).Pluck("id", &wardIds)

	for i := 0; i < 20; i++ {
		stores = append(stores, &models.Store{
			Name:    gofakeit.Company(),
			Phone:   gofakeit.Phone(),
			Address: gofakeit.Address().Address,
			WardId:  gofakeit.RandomUint(wardIds),
			Owner:   gofakeit.Name(),
		})
	}

	database.DB.Create(stores)
	fmt.Println("Seeding Store Successfully")
}
