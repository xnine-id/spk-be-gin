package regency

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findRegencyQs) (*pagination.Pagination[models.Regency], error)
	findById(id string) (*models.Regency, error)
	create(body *regencyBody) (*models.Regency, error)
	update(body *regencyBody, id string) error
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

func (s *service) find(qs *findRegencyQs) (*pagination.Pagination[models.Regency], error) {
	result := pagination.New(models.Regency{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Kabupaten/Kota")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Regency, error) {
	var regency models.Regency

	if err := s.repo.findById(&regency, id); err != nil {
		return nil, customerror.GormError(err, "Kabupaten/Kota")
	}

	return &regency, nil
}

func (s *service) create(body *regencyBody) (*models.Regency, error) {
	regency := models.Regency{
		Name:       body.Name,
		ProvinceId: body.ProvinceId,
	}

	if err := s.repo.create(&regency); err != nil {
		return nil, customerror.GormError(err, "Kabupaten/Kota")
	}

	return &regency, nil
}

func (s *service) update(body *regencyBody, id string) error {
	regency := models.Regency{
		Name:       body.Name,
		ProvinceId: body.ProvinceId,
	}

	if err := s.repo.update(&regency, id); err != nil {
		return customerror.GormError(err, "Kabupaten/Kota")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Kabupaten/Kota")
	}
	return nil
}
