package main

import (
	"github.com/amuhajirs/gin-gorm/src/config"
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/amuhajirs/gin-gorm/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	// helpers.LoadEnvVariables()
	database.ConnectDB()
}

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = config.App.MaxMultipartMemory
	router.Use(middleware.Cors())
	router.Use(middleware.RateLimiter())
	router.Use(middleware.Gzip())

	customValidator := helpers.NewCustomValidator()
	binding.Validator = customValidator

	router.Static("/public", "./public")
	routes.Init(router.Group("/api/v1"))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Rute tidak ditemukan"})
	})

	router.Run()
}
