package store

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Store], query *findStoreQs) error
	findById(store *models.Store, id string) error
	create(store *models.Store) error
	update(store *models.Store, id string) error
	delete(id string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Store], qs *findStoreQs) error {
	q := database.DB.Where("name ILIKE ?", "%"+qs.Search+"%").Preload("Owner")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Order,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(store *models.Store, id string) error {
	return database.DB.Where("id = ?", id).Preload("Owner").First(store).Error
}

func (r *repository) create(store *models.Store) error {
	return database.DB.Create(store).Error
}

func (r *repository) update(store *models.Store, id string) error {
	tx := database.DB.Model(&models.Store{}).Where("id = ?", id).Updates(store)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Toko tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Store{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Toko tidak ditemukan", 404)
	}

	return nil
}
