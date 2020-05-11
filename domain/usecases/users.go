package usecases

import (
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
)

// GetUsers usecase returns user list
func GetUsers(userRepo repositories.IUserRepository) repositories.Users {
	return userRepo.FetchUsers()
}
