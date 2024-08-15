package main

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/database/seeders"
	"github.com/amuhajirs/gin-gorm/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	seeders.UserSeeder()
}
