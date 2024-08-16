package ward

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Ward], query *findWardQs) error
	findById(ward *models.Ward, id string) error
	create(ward *models.Ward) error
	update(ward *models.Ward, id string) error
	delete(id string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Ward], qs *findWardQs) error {
	q := database.DB.Where("name ILIKE ?", "%"+qs.Search+"%").Preload("Subdistrict.Regency.Province")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Order,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(ward *models.Ward, id string) error {
	return database.DB.Where("id = ?", id).Preload("Subdistrict").First(ward).Error
}

func (r *repository) create(ward *models.Ward) error {
	return database.DB.Create(ward).Error
}

func (r *repository) update(ward *models.Ward, id string) error {
	tx := database.DB.Model(&models.Ward{}).Where("id = ?", id).Updates(ward)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kelurahan tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Ward{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kelurahan tidak ditemukan", 404)
	}

	return nil
}
