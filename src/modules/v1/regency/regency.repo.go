package regency

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Regency], query *findRegencyQs) error
	findById(regency *models.Regency, id string) error
	create(regency *models.Regency) error
	update(regency *models.Regency, id string) error
	delete(id string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Regency], qs *findRegencyQs) error {
	q := database.DB.
		Table("mst_regencies mr").
		Where("mr.name ILIKE ?", "%"+qs.Search+"%").
		Joins("LEFT JOIN mst_provinces mp ON mp.id = mr.province_id")
	
	if qs.ProvinceId != "" {
		q = q.Where("mp.id = ?", qs.ProvinceId)
	}

	switch qs.Sort {
	case "name":
		qs.Sort = "mr.name"
	case "province":
		qs.Sort = "mp.name"
	}
	
	q = q.Preload("Province")

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Sort,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(regency *models.Regency, id string) error {
	return database.DB.Where("id = ?", id).Preload("Province").First(regency).Error
}

func (r *repository) create(regency *models.Regency) error {
	return database.DB.Create(regency).Error
}

func (r *repository) update(regency *models.Regency, id string) error {
	tx := database.DB.Model(&models.Regency{}).Where("id = ?", id).Updates(regency)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kabupaten/Kota tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Regency{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kabupaten/Kota tidak ditemukan", 404)
	}

	return nil
}
