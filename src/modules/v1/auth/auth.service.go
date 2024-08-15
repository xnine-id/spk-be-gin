package auth

import (
	"errors"
	"os"
	"time"

	"github.com/amuhajirs/gin-gorm/src/config"
	"github.com/amuhajirs/gin-gorm/src/helpers"
	"github.com/amuhajirs/gin-gorm/src/helpers/customerror"
	"github.com/amuhajirs/gin-gorm/src/helpers/jwt"
	"github.com/amuhajirs/gin-gorm/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	login(body *loginBody) (*models.User, *userToken, error)
	refresh(body *refreshBody) (*userToken, error)
	logout(body *refreshBody) error
	generateUserToken(id uint) *userToken
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) login(body *loginBody) (*models.User, *userToken, error) {
	var user models.User

	if err := s.repo.findByUsername(&user, body.Username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, customerror.New("Nama pengguna atau kata sandi salah", 400)
		}
		return nil, nil, customerror.New("Terjadi kesalahan saat mengambil data", 500)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return nil, nil, customerror.New("Nama pengguna atau kata sandi salah", 400)
	}

	token := s.generateUserToken(*user.Id)

	go s.repo.createToken(&models.Token{
		UserId:   *user.Id,
		Token:    token.RefreshToken,
		LastUsed: time.Now(),
	})

	user.Password = ""

	return &user, token, nil
}

func (s *service) refresh(body *refreshBody) (*userToken, error) {
	var refreshToken models.Token

	if err := s.repo.findToken(&refreshToken, body.RefreshToken); err != nil {
		return nil, customerror.New("Refresh token tidak valid", 400)
	}

	if refreshToken.User == nil {
		return nil, customerror.New("Pengguna tidak ditemukan", 400)
	}

	// if time.Since(refreshToken.LastUsed) < config.App.Auth.AccessTokenExpiresIn {
	// 	return nil, customerror.New("Belum bisa refresh token", 400)
	// }

	if time.Since(refreshToken.LastUsed) > config.App.Auth.RefreshTokenExpiresIn {
		go s.repo.deleteTokenByToken(refreshToken.Token)
		return nil, customerror.New("Refresh token kadaluarsa", 400)
	}

	refreshToken.LastUsed = time.Now()
	token := s.generateUserToken(*refreshToken.User.Id)
	refreshToken.Token = token.RefreshToken

	go s.repo.updateToken(&refreshToken)

	return token, nil
}

func (s *service) logout(body *refreshBody) error {
	return s.repo.deleteTokenByToken(body.RefreshToken)
}

func (s *service) generateUserToken(id uint) *userToken {
	var token userToken

	token.AccessToken = jwt.New(&jwt.MapClaims{
		"iss":  os.Getenv("BASE_URL"),
		"sub":  id,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(config.App.Auth.AccessTokenExpiresIn).Unix(),
		"type": "access",
	})

	token.RefreshToken = helpers.RandomString()
	token.Type = "Bearer"
	token.ExpiresIn = config.App.Auth.AccessTokenExpiresIn / time.Second
	token.RefreshTokenExpiresIn = config.App.Auth.RefreshTokenExpiresIn / time.Second

	return &token
}
