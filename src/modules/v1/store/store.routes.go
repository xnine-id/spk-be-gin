package store

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	storeRepo       = NewRepository()
	storeService    = NewService(storeRepo)
	storeController = NewController(storeService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("", storeController.find)
		authorized.GET("/:id", storeController.findById)
		authorized.POST("", storeController.create)
		authorized.PUT("/:id", storeController.update)
		authorized.DELETE("/:id", storeController.delete)
	}
}
