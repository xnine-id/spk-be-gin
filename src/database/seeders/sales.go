package seeders

import (
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/brianvoe/gofakeit/v7"
)

func SalesSeeder() {
	var sales []*models.Sales
	var wardIds []uint

	database.DB.Model(models.Ward{}).Pluck("id", &wardIds)

	for i := 0; i < 20; i++ {
		sales = append(sales, &models.Sales{
			Name: gofakeit.Name(),
			Email: gofakeit.Email(),
			Phone: gofakeit.Phone(),
			Address: gofakeit.Address().Address,
			WardId: gofakeit.RandomUint(wardIds),
			Photo: "/public/assets/avatar-user.jpg",
		})
	}

	database.DB.Create(sales)
	fmt.Println("Seeding Sales Successfully")
}