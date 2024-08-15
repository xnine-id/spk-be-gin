package seeders

import (
	"fmt"
	"os"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/brianvoe/gofakeit/v7"
	"golang.org/x/crypto/bcrypt"
)

func UserSeeder() {
	var users []*models.User

	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	avatar := os.Getenv("BASE_URL") + "/public/assets/avatar-user.jpg"

	users = append(users, &models.User{
		Name:     "Developer",
		Username: "dev",
		Phone:    helpers.PointerTo("08" + gofakeit.Numerify("##########")),
		Password: string(password),
		Avatar:   avatar,
	})

	for i := 0; i < 9; i++ {
		users = append(users, &models.User{
			Name:     gofakeit.Name(),
			Username: gofakeit.Username(),
			Phone:    helpers.PointerTo("08" + gofakeit.Numerify("##########")),
			Password: string(password),
			Avatar:   avatar,
		})
	}

	database.DB.Create(users)

	fmt.Println("User Seeder created")
}
