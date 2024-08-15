package ward

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findWardQs) (*pagination.Pagination[models.Ward], error)
	findById(id string) (*models.Ward, error)
	create(body *wardBody) (*models.Ward, error)
	update(body *wardBody, id string) error
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

func (s *service) find(qs *findWardQs) (*pagination.Pagination[models.Ward], error) {
	result := pagination.New(models.Ward{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Kelurahan")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Ward, error) {
	var ward models.Ward

	if err := s.repo.findById(&ward, id); err != nil {
		return nil, customerror.GormError(err, "Kelurahan")
	}

	return &ward, nil
}

func (s *service) create(body *wardBody) (*models.Ward, error) {
	ward := models.Ward{
		Name:          body.Name,
		SubdistrictId: body.SubdistrictId,
	}

	if err := s.repo.create(&ward); err != nil {
		return nil, customerror.GormError(err, "Kelurahan")
	}

	return &ward, nil
}

func (s *service) update(body *wardBody, id string) error {
	ward := models.Ward{
		Name:          body.Name,
		SubdistrictId: body.SubdistrictId,
	}

	if err := s.repo.update(&ward, id); err != nil {
		return customerror.GormError(err, "Kelurahan")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Kelurahan")
	}
	return nil
}
