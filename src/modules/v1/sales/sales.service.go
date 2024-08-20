package sales

import (
	"strconv"

	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/helpers/upload"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findSalesQs) (*pagination.Pagination[models.Sales], error)
	findById(id string) (*models.Sales, error)
	create(body *createSalesBody) (*models.Sales, error)
	update(body *updateSalesBody, id string) error
	delete(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) find(qs *findSalesQs) (*pagination.Pagination[models.Sales], error) {
	result := pagination.New(models.Sales{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Sales")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Sales, error) {
	var sales models.Sales

	if err := s.repo.findById(&sales, id); err != nil {
		return nil, customerror.GormError(err, "Sales")
	}

	return &sales, nil
}

func (s *service) create(body *createSalesBody) (*models.Sales, error) {
	sales := models.Sales{
		Name:    body.Name,
		Email:   body.Email,
		Phone:   body.Phone,
		Address: body.Address,
		WardId:  body.WardId,
	}

	if err := s.repo.create(&sales); err != nil {
		return nil, customerror.GormError(err, "Sales")
	}

	filePath, err := upload.New(&upload.Option{
		Folder:      "sales",
		File:        body.Photo,
		NewFilename: strconv.FormatUint(uint64(*sales.Id), 10),
	})

	if err != nil {
		return nil, customerror.New("Gagal saat menyimpan file", 500)
	}

	sales.Photo = filePath.Url

	if err := s.repo.save(&sales); err != nil {
		return nil, customerror.GormError(err, "Sales")
	}

	return &sales, nil
}

func (s *service) update(body *updateSalesBody, id string) error {
	var sales models.Sales

	if err := s.repo.findById(&sales, id); err != nil {
		return customerror.GormError(err, "Sales")
	}

	if body.Photo != nil {
		filePath, err := upload.New(&upload.Option{
			Folder:      "sales",
			File:        body.Photo,
			NewFilename: strconv.FormatUint(uint64(*sales.Id), 10),
		})
		
		if err != nil {
			return customerror.New("Gagal saat menyimpan file", 500)
		}

		sales.Photo = filePath.Path
	}


	sales.Name = body.Name
	sales.Email = body.Email
	sales.Phone = body.Phone
	sales.Address = body.Address
	sales.WardId = body.WardId

	if err := s.repo.save(&sales); err != nil {
		return customerror.GormError(err, "Sales")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Sales")
	}
	return nil
}
