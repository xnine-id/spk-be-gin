package ward

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	wardRepo       = NewRepository()
	wardService    = NewService(wardRepo)
	wardController = NewController(wardService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/", wardController.find)
		authorized.GET("/:id", wardController.findById)
		authorized.POST("/", wardController.create)
		authorized.PUT("/:id", wardController.update)
		authorized.DELETE("/:id", wardController.delete)
	}
}
