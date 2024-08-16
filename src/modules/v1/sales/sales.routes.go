package sales

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	salesRepo       = NewRepository()
	salesService    = NewService(salesRepo)
	salesController = NewController(salesService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("", salesController.find)
		authorized.GET("/:id", salesController.findById)
		authorized.POST("", salesController.create)
		authorized.PUT("/:id", salesController.update)
		authorized.DELETE("/:id", salesController.delete)
	}
}
