package response

import (
	"errors"
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/gin-gonic/gin"
)

func ServiceError(ctx *gin.Context, err error) {
	if err != nil {
		var customErr *customerror.CustomError
		if errors.As(err, &customErr) {
			Error(ctx, customErr.Message, customErr.StatusCode)
		} else {
			fmt.Println(err.Error())
			Error(ctx, "Terjadi kesalahan", 500)
		}
	}
}

func Error(ctx *gin.Context, msg string, statusCode int) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{"message": msg})
}