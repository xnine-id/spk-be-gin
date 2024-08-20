package pagination

import (
	"math"
	"strconv"

	"github.com/amuhajirs/gin-gorm/src/database"
	"gorm.io/gorm"
)

type meta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	LastPage    int   `json:"last_page"`
}

type Pagination[T any] struct {
	Meta meta `json:"meta"`
	Data *[]T `json:"data"`
}

type Params struct {
	Query     *gorm.DB
	Page      string
	Limit     string
	Order     string
	Direction string
}

type QS struct {
	Page      string `form:"page"`
	Search    string `form:"search" mod:"trim"`
	Limit     string `form:"limit"`
	Sort      string `form:"sort"`
	Direction string `form:"direction"`
}

func New[T any](model T) *Pagination[T] {
	return &Pagination[T]{Data: &[]T{}}
}

func (pagination *Pagination[any]) Execute(params *Params) error {
	if params.Query == nil {
		params.Query = database.DB
	}

	if params.Page == "" {
		pagination.Meta.CurrentPage = 1
	} else {
		pagination.Meta.CurrentPage, _ = strconv.Atoi(params.Page)
	}

	if params.Limit == "" {
		pagination.Meta.PerPage = 10
	} else {
		pagination.Meta.PerPage, _ = strconv.Atoi(params.Limit)
	}

	if params.Order == "" {
		params.Order = "created_at"
	}

	if params.Direction == "" {
		params.Direction = "desc"
	}

	var totalData int64
	params.Query.Model(pagination.Data).Count(&totalData)
	offset := (pagination.Meta.CurrentPage - 1) * pagination.Meta.PerPage

	pagination.Meta.Total = totalData
	pagination.Meta.LastPage = int(math.Ceil(float64(totalData) / float64(pagination.Meta.PerPage)))

	return params.Query.Offset(offset).Limit(pagination.Meta.PerPage).Order(params.Order + " " + params.Direction).Find(pagination.Data).Error
}
