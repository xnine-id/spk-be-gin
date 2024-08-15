package middleware

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/amuhajirs/gin-gorm/src/config"
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")

		allowedOrigin := helpers.Find(&config.App.Cors.AllowOrigin, func(t *string) bool {
			return *t == origin || *t == "*"
		})

		if allowedOrigin != nil {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", *allowedOrigin)
		}
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", strconv.FormatBool(config.App.Cors.AllowCredential))
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(config.App.Cors.AllowHeaders, ", "))
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(config.App.Cors.AllowMethod, ", "))

		if ctx.Request.Method == "OPTIONS" {
			fmt.Println(*allowedOrigin, ctx.Request.URL.Path, origin)
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
