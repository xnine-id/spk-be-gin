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
	q := database.DB.
		Table("mst_wards mw").
		Joins("LEFT JOIN mst_subdistricts ms ON ms.id = mw.subdistrict_id").
		Joins("LEFT JOIN mst_regencies mr ON mr.id = ms.regency_id").
		Joins("LEFT JOIN mst_provinces mp ON mp.id = mr.province_id").
		Where("mw.name ILIKE ?", "%"+qs.Search+"%").
		Preload("Subdistrict.Regency.Province")

	if qs.ProvinceId != "" {
		q = q.Where("mp.id = ?", qs.ProvinceId)
	}

	if qs.RegencyId != "" {
		q = q.Where("mr.id = ?", qs.RegencyId)
	}

	if qs.SubdistrictId != "" {
		q = q.Where("ms.id = ?", qs.SubdistrictId)
	}

	switch qs.Sort {
	case "name":
		qs.Sort = "mw.name"
	case "subdistrict":
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

func (r *repository) findById(ward *models.Ward, id string) error {
	return database.DB.Where("id = ?", id).Preload("Subdistrict.Regency.Province").First(ward).Error
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
