package helpers

import (
	"github.com/gin-gonic/gin"
)

func Bind(ctx *gin.Context, model interface{}) (isValid bool) {
	if err := ctx.ShouldBind(model); err != nil {
		parsed, parseErr := ParseError(err)

		if parseErr == nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"errors":  &parsed,
				"message": parsed[0].Msg,
			})

			return
		} else {
			ctx.AbortWithStatusJSON(500, gin.H{
				"message": err.Error(),
			})

			return
		}
	}

	return true
}
