package subdistrict

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	subdistrictRepo       = NewRepository()
	subdistrictService    = NewService(subdistrictRepo)
	subdistrictController = NewController(subdistrictService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("", subdistrictController.find)
		authorized.GET("/:id", subdistrictController.findById)
		authorized.POST("", subdistrictController.create)
		authorized.PUT("/:id", subdistrictController.update)
		authorized.DELETE("/:id", subdistrictController.delete)
	}
}
