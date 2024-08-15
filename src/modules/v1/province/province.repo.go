package province

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Province], query *findProvinceQs) error
	findById(province *models.Province, id string) error
	create(province *models.Province) error
	update(province *models.Province, id string) error
	delete(id string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Province], qs *findProvinceQs) error {
	q := database.DB.Where("name ILIKE ?", "%"+qs.Search+"%")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Order,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(province *models.Province, id string) error {
	return database.DB.Where("id = ?", id).First(province).Error
}

func (r *repository) create(province *models.Province) error {
	return database.DB.Create(province).Error
}

func (r *repository) update(province *models.Province, id string) error {
	tx := database.DB.Model(&models.Province{}).Where("id = ?", id).Updates(province)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Provinsi tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Province{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Provinsi tidak ditemukan", 404)
	}

	return nil
}
