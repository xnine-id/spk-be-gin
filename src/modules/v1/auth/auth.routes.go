package auth

import (
	"github.com/amuhajirs/gin-gorm/src/middleware"
	"github.com/gin-gonic/gin"
)

var (
	authRepo       = NewRepository()
	authService    = NewService(authRepo)
	authController = NewController(authService)
)

func Routes(router *gin.RouterGroup) {
	router.POST("/login", authController.login)
	router.POST("/refresh", authController.refresh)
	router.POST("/logout", authController.logout)

	authorized := router.Group("")
	authorized.Use(middleware.Auth())
	{
		authorized.GET("/profile", authController.profile)
	}
}
