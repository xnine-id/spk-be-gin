package auth

import (
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/amuhajirs/gin-gorm/src/helpers/response"
	"github.com/amuhajirs/gin-gorm/src/models"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	login(ctx *gin.Context)
	profile(ctx *gin.Context)
	refresh(ctx *gin.Context)
	logout(ctx *gin.Context)
}

type controller struct {
	service Service
}

func NewController(service Service) Controller {
	return &controller{
		service: service,
	}
}

func (c *controller) login(ctx *gin.Context) {
	var body loginBody

	if isValid := helpers.Bind(ctx, &body); !isValid {
		return
	}

	data, token, err := c.service.login(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Login berhasil", "token": token, "data": data})
}

func (c *controller) refresh(ctx *gin.Context) {
	var body refreshBody

	if isValid := helpers.Bind(ctx, &body); !isValid {
		return
	}

	token, err := c.service.refresh(&body)

	if err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func (c *controller) profile(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	_user := user.(models.User)
	_user.Password = ""

	ctx.JSON(200, gin.H{
		"data": _user,
	})
}

func (c *controller) logout(ctx *gin.Context) {
	var body refreshBody

	if isValid := helpers.Bind(ctx, &body); !isValid {
		return
	}

	if err := c.service.logout(&body); err != nil {
		response.ServiceError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Logout berhasil",
	})
}
