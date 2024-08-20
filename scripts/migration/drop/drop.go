package main

import (
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers"
)

func init() {
	helpers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	tables, _ := database.DB.Migrator().GetTables()

	_tables := []interface{}{}
	for _, v := range tables {
		_tables = append(_tables, v)
	}
	database.DB.Migrator().DropTable(_tables...)

	fmt.Println("Tables Dropped")
}
