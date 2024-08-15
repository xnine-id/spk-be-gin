package regency

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	regencyRepo       = NewRepository()
	regencyService    = NewService(regencyRepo)
	regencyController = NewController(regencyService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("", regencyController.find)
		authorized.GET("/:id", regencyController.findById)
		authorized.POST("", regencyController.create)
		authorized.PUT("/:id", regencyController.update)
		authorized.DELETE("/:id", regencyController.delete)
	}
}
