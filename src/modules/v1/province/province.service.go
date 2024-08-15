package province

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findProvinceQs) (*pagination.Pagination[models.Province], error)
	findById(id string) (*models.Province, error)
	create(body *provinceBody) (*models.Province, error)
	update(body *provinceBody, id string) error
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

func (s *service) find(qs *findProvinceQs) (*pagination.Pagination[models.Province], error) {
	result := pagination.New(models.Province{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Provinsi")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Province, error) {
	var province models.Province

	if err := s.repo.findById(&province, id); err != nil {
		return nil, customerror.GormError(err, "Provinsi")
	}

	return &province, nil
}

func (s *service) create(body *provinceBody) (*models.Province, error) {
	province := models.Province{
		Name: body.Name,
	}

	if err := s.repo.create(&province); err != nil {
		return nil, customerror.GormError(err, "Provinsi")
	}

	return &province, nil
}

func (s *service) update(body *provinceBody, id string) error {
	province := models.Province{
		Name: body.Name,
	}

	if err := s.repo.update(&province, id); err != nil {
		return customerror.GormError(err, "Provinsi")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Provinsi")
	}
	return nil
}
