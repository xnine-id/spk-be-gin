package subdistrict

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	find(result *pagination.Pagination[models.Subdistrict], query *findSubdistrictQs) error
	findById(subdistrict *models.Subdistrict, id string) error
	create(subdistrict *models.Subdistrict) error
	update(subdistrict *models.Subdistrict, id string) error
	delete(id string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Subdistrict], qs *findSubdistrictQs) error {
	q := database.DB.
		Table("mst_subdistricts ms").
		Joins("LEFT JOIN mst_regencies mr ON mr.id = ms.regency_id").
		Joins("LEFT JOIN mst_provinces mp ON mp.id = mr.province_id").
		Where("ms.name ILIKE ?", "%"+qs.Search+"%").
		Preload("Regency.Province")

	if qs.ProvinceId != "" {
		q = q.Where("mp.id = ?", qs.ProvinceId)
	}

	if qs.RegencyId != "" {
		q = q.Where("mr.id = ?", qs.RegencyId)
	}

	switch qs.Sort {
	case "name":
		qs.Sort = "ms.name"
	case "regency":
		qs.Sort = "mr.name"
	case "province":
		qs.Sort = "mp.name"
	}

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Sort,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(subdistrict *models.Subdistrict, id string) error {
	return database.DB.Where("id = ?", id).Preload("Regency.Province").First(subdistrict).Error
}

func (r *repository) create(subdistrict *models.Subdistrict) error {
	return database.DB.Create(subdistrict).Error
}

func (r *repository) update(subdistrict *models.Subdistrict, id string) error {
	tx := database.DB.Model(&models.Subdistrict{}).Where("id = ?", id).Updates(subdistrict)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kecamatan tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Subdistrict{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Kecamatan tidak ditemukan", 404)
	}

	return nil
}
