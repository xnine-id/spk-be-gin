package installation

import (
	"errors"
	"fmt"

	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
	"gorm.io/gorm"
)

type Repository interface {
	find(result *pagination.Pagination[models.Installation], query *findInstallationQs) error
	findById(installation *models.Installation, id string) error
	save(installation *models.Installation) error
	delete(id string) error
	importData(installations *[]models.Installation) error

	saveImages(installationImgs *[]models.InstallationImage) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) find(result *pagination.Pagination[models.Installation], qs *findInstallationQs) error {
	q := database.DB.
		Table("trx_installations ti").
		Joins("LEFT JOIN mst_sales msa ON msa.id = ti.sales_id").
		Joins("LEFT JOIN mst_stores mst ON mst.id = ti.store_id").
		Where("ti.spk_number ILIKE ?", "%"+qs.Search+"%").
		Preload("Store").
		Preload("Sales").
		Preload("Images")

	switch qs.Sort {
	case "sales":
		qs.Sort = "msa.name"
	case "store":
		qs.Sort = "mst.name"
	}

	return result.Execute(&pagination.Params{
		Query:     q,
		Page:      qs.Page,
		Limit:     qs.Limit,
		Order:     qs.Sort,
		Direction: qs.Direction,
	})
}

func (r *repository) findById(installation *models.Installation, id string) error {
	return database.DB.
		Where("id = ?", id).
		Preload("Store").
		Preload("Sales").
		Preload("Images").
		First(installation).Error
}

func (r *repository) save(installation *models.Installation) error {
	return database.DB.Save(installation).Error
}

func (r *repository) delete(id string) error {
	tx := database.DB.Where("id = ?", id).Delete(&models.Installation{})

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return customerror.New("Pemasangan tidak ditemukan", 404)
	}

	return nil
}

func (r *repository) importData(installations *[]models.Installation) error {
	return  database.DB.Transaction(func(tx *gorm.DB) error {
		// Iterate over each installation and attempt to create it in the database
		for i, inst := range *installations {
			if err := tx.Select("spk_number", "spk_date", "store_id").Create(&inst).Error; err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return customerror.New(fmt.Sprintf("Terdapat key yang duplikat pada baris ke %d", i + 1), 400)
				}
			
				if errors.Is(err, gorm.ErrForeignKeyViolated) {
					return customerror.New(fmt.Sprintf("Foreign key tidak valid pada baris ke %d", i + 1), 400)
				}
			}
		}

		return nil
	})
}

func (r *repository) saveImages(installationImgs *[]models.InstallationImage) error {
	return database.DB.Save(installationImgs).Error
}