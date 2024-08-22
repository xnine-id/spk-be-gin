package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func ParseError(err error, validatedStruct any) []ApiError {
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		reflectedStruct := reflect.ValueOf(validatedStruct)

		for i, fe := range ve {
			fieldName := fe.StructField()

			// Get the field's value from the struct
			field, ok := reflectedStruct.Type().Elem().FieldByName(fieldName)
			if !ok {
				// If field is not found, continue or handle as needed
				continue
			}

			out[i] = ApiError{
				Field: helpers.ToSnakeCase(fe.Field()),
				Msg:   MsgForTag(fe.Tag(), fe.Field(), fe.Param(), field.Type.String()),
			}
		}
		return out
	}

	return nil
}

func ParseUnmarshalError(err error) *string {
	var ue *json.UnmarshalTypeError

	if errors.As(err, &ue) {
		return helpers.PointerTo(fmt.Sprintf("Gagal memproses data, pastikan `%s` berupa `%s`", ue.Field, ue.Type))
	}
	return nil
}
