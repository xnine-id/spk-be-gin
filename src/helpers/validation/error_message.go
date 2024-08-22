package validation

import (
	"fmt"
	"strings"

	"github.com/amuhajirs/gin-gorm/src/config"
)

func fieldToLocale(field string) string {
	for k, v := range config.Locale {
		if field == k {
			return v
		}
	}

	words := strings.Split(field, "_")

	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}

	return strings.Join(words, " ")
}

func MsgForTag(tag string, field string, param string, dataType string) string {
	fmt.Println(dataType)
	intlField := fieldToLocale(field)

	switch tag {
	case "required":
		return intlField + " harus diisi"
	case "email":
		return "Email tidak valid"
	case "min":
		return intlField + " minimal " + param + " karakter"
	case "max":
		return intlField + " maximal " + param + " karakter"
	case "len":
		if dataType == "string" {
			return intlField + " harus " + param + " karakter"
		}

		return "Jumlah file yang diunggah harus " + param
	case "unique":
		return intlField + " sudah digunakan"
	case "numeric":
		return intlField + " harus berupa angka"
	case "required_if":
		return intlField + " harus diisi"
	case "datetime":
		return "Format waktu tidak valid"
	case "oneof":
		enum := strings.Join(strings.Split(param, " "), " atau ")
		return intlField + " harus " + enum
	case "image":
		return intlField + " harus file gambar"
	}
	return ""
}
