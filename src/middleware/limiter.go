package middleware

import (
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/config"
	"github.com/amuhajirs/gin-gorm/src/helpers/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// rate (20/s) and burst (30/s)
var limiter = rate.NewLimiter(rate.Limit(config.App.Limiter.Rate), config.App.Limiter.Burst)

func RateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			response.Error(ctx, "Too Many Request", 429)
			fmt.Println("Too Many Request")
		}
	}
}
