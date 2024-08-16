package province

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/response"
	"github.com/amuhajirs/gin-gorm/src/helpers/validation"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	find(ctx *gin.Context)
	findById(ctx *gin.Context)
	create(ctx *gin.Context)
	update(ctx *gin.Context)
	delete(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) find(ctx *gin.Context) {
	var qs findProvinceQs

	ctx.ShouldBindQuery(&qs)

	data, err := c.service.find(&qs)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, data)
}

func (c *controller) findById(ctx *gin.Context) {
	id := ctx.Param("id")

	data, err := c.service.findById(id)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"data": data})
}

func (c *controller) create(ctx *gin.Context) {
	var body provinceBody

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	data, err := c.service.create(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Provinsi berhasil ditambahkan",
		"data":    data,
	})
}

func (c *controller) update(ctx *gin.Context) {
	var body provinceBody
	id := ctx.Param("id")

	if isValid := validation.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.update(&body, id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Provinsi berhasil diperbarui"})
}

func (c *controller) delete(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := c.service.delete(id); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Provinsi berhasil dihapus"})
}
