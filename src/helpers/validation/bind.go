package validation

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/response"
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) (isValid bool) {
	if err := ctx.ShouldBind(model); err != nil {
		if parsed := ParseError(err); parsed != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"errors":  &parsed,
				"message": parsed[0].Msg,
			})
			return
		}

		if parsed := ParseUnmarshalError(err); parsed != nil {
			response.Error(ctx, *parsed, 500)
			return
		}

		response.Error(ctx, err.Error(), 500)
		return 
	}

	return true
}
