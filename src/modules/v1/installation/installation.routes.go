package installation

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	installationRepo       = NewRepository()
	installationService    = NewService(installationRepo)
	installationController = NewController(installationService)
)

func Routes(router *gin.RouterGroup) {
	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("", installationController.find)
		authorized.GET("/:id", installationController.findById)
		authorized.POST("", installationController.create)
		authorized.PUT("/:id", installationController.update)
		authorized.DELETE("/:id", installationController.delete)
		authorized.POST("/import", installationController.importData)
	}
}
