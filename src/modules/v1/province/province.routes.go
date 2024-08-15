package province

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	provinceRepo       = NewRepository()
	provinceService    = NewService(provinceRepo)
	provinceController = NewController(provinceService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("/")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/", provinceController.find)
		authorized.GET("/:id", provinceController.findById)
		authorized.POST("/", provinceController.create)
		authorized.PUT("/:id", provinceController.update)
		authorized.DELETE("/:id", provinceController.delete)
	}
}
