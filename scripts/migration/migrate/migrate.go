package main

import (
	"fmt"

	"github.com/amuhajirs/gin-gorm/scripts/migration"
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	database.DB.AutoMigrate(migration.Tables...)

	fmt.Println("Tables Created")
}