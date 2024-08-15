package subdistrict

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findSubdistrictQs) (*pagination.Pagination[models.Subdistrict], error)
	findById(id string) (*models.Subdistrict, error)
	create(body *subdistrictBody) (*models.Subdistrict, error)
	update(body *subdistrictBody, id string) error
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

func (s *service) find(qs *findSubdistrictQs) (*pagination.Pagination[models.Subdistrict], error) {
	result := pagination.New(models.Subdistrict{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Kecamatan")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Subdistrict, error) {
	var subdistrict models.Subdistrict

	if err := s.repo.findById(&subdistrict, id); err != nil {
		return nil, customerror.GormError(err, "Kecamatan")
	}

	return &subdistrict, nil
}

func (s *service) create(body *subdistrictBody) (*models.Subdistrict, error) {
	subdistrict := models.Subdistrict{
		Name:      body.Name,
		RegencyId: body.RegencyId,
	}

	if err := s.repo.create(&subdistrict); err != nil {
		return nil, customerror.GormError(err, "Kecamatan")
	}

	return &subdistrict, nil
}

func (s *service) update(body *subdistrictBody, id string) error {
	subdistrict := models.Subdistrict{
		Name:      body.Name,
		RegencyId: body.RegencyId,
	}

	if err := s.repo.update(&subdistrict, id); err != nil {
		return customerror.GormError(err, "Kecamatan")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Kecamatan")
	}
	return nil
}
