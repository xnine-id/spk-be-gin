package routes

import (
	"github.com/amuhajirs/gin-gorm/src/modules/v1/auth"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/province"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/regency"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/sales"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/store"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/subdistrict"
	"github.com/amuhajirs/gin-gorm/src/modules/v1/ward"
	"github.com/gin-gonic/gin"
)

func Init(router *gin.RouterGroup) {
	auth.Routes(router.Group("/auth"))
	province.Routes(router.Group("/provinces"))
	regency.Routes(router.Group("/regencies"))
	subdistrict.Routes(router.Group("/subdistricts"))
	ward.Routes(router.Group("/wards"))
	store.Routes(router.Group("/stores"))
	sales.Routes(router.Group("/sales"))
}
