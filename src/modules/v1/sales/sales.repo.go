package sales

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Sales], query *findSalesQs) error
	findById(sales *models.Sales, id string) error
	create(sales *models.Sales) error
	update(sales *models.Sales, id string) error
	delete(id string) error
	save(sales *models.Sales) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Sales], qs *findSalesQs) error {
	q := database.DB.Where("name ILIKE ?", "%"+qs.Search+"%")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Sort,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(sales *models.Sales, id string) error {
	return database.DB.Where("id = ?", id).First(sales).Error
}

func (r *repository) create(sales *models.Sales) error {
	return database.DB.Create(sales).Error
}

func (r *repository) update(sales *models.Sales, id string) error {
	tx := database.DB.Model(&models.Sales{}).Where("id = ?", id).Updates(sales)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Sales tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Sales{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Sales tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) save(sales *models.Sales) error {
	tx := database.DB.Save(sales)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Sales tidak ditemukan", 404)
	}

	return nil
}
