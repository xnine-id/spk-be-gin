package helpers

import (
	crRand "crypto/rand"
	"encoding/base64"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func PointerTo[T any](s T) *T {
	return &s
}

func RandomElement[T any](nums []T) T {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	return nums[s.Intn(len(nums))]
}

func Find[T any](slice *[]T, test func(*T) bool) (ret *T) {
	for _, s := range *slice {
		if test(&s) {
			ret = &s
		}
	}
	return
}

func RandomString() string {
	bytes := make([]byte, 40)
	crRand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes)
}