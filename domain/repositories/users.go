package repositories

import (
	"errors"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/jinzhu/gorm"
)

// UserRepository ..
type UserRepository struct {
	db *gorm.DB
}

// Users is list of users
type Users []models.User

// IUserRepository is a interface of UserRepository
type IUserRepository interface {
	FetchUsers() Users
	FetchUserByID(userID uint) (user models.User, err error)
	FetchUserByUsername(username string) (user models.User, err error)
}

// NewUserRepo returns new UserRepository
func NewUserRepo(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// FetchUsers Repository returns user list
func (r *UserRepository) FetchUsers() Users {
	var users Users
	r.db.Find(&users)

	return users
}

// FetchUserByID repository returns user
func (r *UserRepository) FetchUserByID(userID uint) (user models.User, err error) {
	result := r.db.First(&user, "id = ?", userID)

	if result.Error != nil {
		err = errors.New("Invalid credentials")
	}

	return
}

// FetchUserByUsername repository returns user
func (r *UserRepository) FetchUserByUsername(username string) (user models.User, err error) {
	result := r.db.Where(&models.User{Username: username}).First(&user)

	if result.Error != nil {
		err = errors.New("Invalid credentials")
	}

	return
}
