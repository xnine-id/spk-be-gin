package auth

import (
	"github.com/amuhajirs/gin-gorm/src/database"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/models"
)

type Repository interface {
	findByUsername(user *models.User, username string) error
	findById(user *models.User, id uint) error

	findToken(model *models.Token, token string) error
	createToken(data *models.Token) error
	updateToken(data *models.Token) error
	deleteTokenByToken(token string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) findByUsername(user *models.User, username string) error {
	return database.DB.Where("username = ?", username).First(user).Error
}

func (r *repository) findById(user *models.User, id uint) error {
	return database.DB.Where("id = ?", id).First(user).Error
}

func (r *repository) findToken(model *models.Token, token string) error {
	return database.DB.Where("token = ?", token).Preload("User").First(model).Error
}

func (r *repository) createToken(data *models.Token) error {
	return database.DB.Create(data).Error
}

func (r *repository) updateToken(data *models.Token) error {
	return database.DB.Save(data).Error
}

func (r *repository) deleteTokenByToken(token string) error {
	if affected := database.DB.Unscoped().Where("token = ?", token).Delete(&models.Token{}).RowsAffected; affected == 0 {
		return customerror.New("Token tidak ditemukan", 404)
	}

	return nil
}
