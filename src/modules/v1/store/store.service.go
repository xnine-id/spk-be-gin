package store

import (
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/pagination"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Service interface {
	find(qs *findStoreQs) (*pagination.Pagination[models.Store], error)
	findById(id string) (*models.Store, error)
	create(body *storeBody) (*models.Store, error)
	update(body *storeBody, id string) error
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

func (s *service) find(qs *findStoreQs) (*pagination.Pagination[models.Store], error) {
	result := pagination.New(models.Store{})

	if err := s.repo.find(result, qs); err != nil {
		return nil, customerror.GormError(err, "Toko")
	}

	return result, nil
}

func (s *service) findById(id string) (*models.Store, error) {
	var store models.Store

	if err := s.repo.findById(&store, id); err != nil {
		return nil, customerror.GormError(err, "Toko")
	}

	return &store, nil
}

func (s *service) create(body *storeBody) (*models.Store, error) {
	store := models.Store{
		Name: body.Name,
		Phone: body.Phone,
		Address: body.Address,
		WardId: body.WardId,
		OwnerId: body.OwnerId,
	}

	if err := s.repo.create(&store); err != nil {
		return nil, customerror.GormError(err, "Toko")
	}

	return &store, nil
}

func (s *service) update(body *storeBody, id string) error {
	store := models.Store{
		Name: body.Name,
		Phone: body.Phone,
		Address: body.Address,
		WardId: body.WardId,
		OwnerId: body.OwnerId,
	}

	if err := s.repo.update(&store, id); err != nil {
		return customerror.GormError(err, "Toko")
	}

	return nil
}

func (s *service) delete(id string) error {
	if err := s.repo.delete(id); err != nil {
		return customerror.GormError(err, "Toko")
	}
	return nil
}
