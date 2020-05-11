package repositories

import (
	"errors"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/jinzhu/gorm"
)

// AuthRepository ..
type AuthRepository struct {
	db *gorm.DB
}

// IAuthRepository is a interface of AuthRepository
type IAuthRepository interface {
	CreateAccessToken(accessToken string, user models.User) models.Accesstoken
	GetUser(username string, password string) (user models.User, err error)
	CreateUser(username string, password string, fullname string) models.User
	CheckAccessToken(token string) (accessToken models.Accesstoken, err error)
}

// NewAuthRepo ..
func NewAuthRepo(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

// CreateAccessToken repository returns auth access Token
func (r *AuthRepository) CreateAccessToken(accessToken string, user models.User) models.Accesstoken {
	var token = models.Accesstoken{
		Token:  accessToken,
		UserID: user.ID,
	}
	r.db.Create(&token)

	return token
}

// GetUser returns specific user
func (r *AuthRepository) GetUser(username string, password string) (user models.User, err error) {
	result := r.db.Where(&models.User{Username: username, Password: password}).First(&user)

	if result.Error != nil {
		err = errors.New("Invalid credentials")
	}

	return
}

// CreateUser repository creates user
func (r *AuthRepository) CreateUser(username string, password string, fullname string) models.User {
	var user = models.User{
		Username: username,
		Password: password,
		FullName: fullname,
	}
	r.db.Create(&user)

	return user
}

// CheckAccessToken checks Access Token
func (r *AuthRepository) CheckAccessToken(token string) (accessToken models.Accesstoken, err error) {
	result := r.db.Where(&models.Accesstoken{Token: token}).First(&accessToken)

	if result.Error != nil {
		err = errors.New("Invalid credentials")
	}

	return
}
