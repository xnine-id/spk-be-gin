package middleware

import (
	"errors"
	"strings"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/jwt"
	"github.com/amuhajirs/gin-gorm/src/helpers/response"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")

		if token == "" {
			response.Error(ctx, "Anda belum login!", 401)
			return
		}

		token = strings.Split(token, " ")[1]

		claim, err := jwt.Parse(token)
		if err != nil {
			response.Error(ctx, "Token tidak valid", 401)
			return
		}

		if claim["type"].(string) != "access" {
			response.Error(ctx, "Token tidak valid", 401)
			return
		}

		var user models.User
		if err := database.DB.Where("id = ?", claim["sub"]).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				response.Error(ctx, "Pengguna tidak ditemukan", 401)
			} else {
				response.Error(ctx, err.Error(), 500)
			}
			return
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
