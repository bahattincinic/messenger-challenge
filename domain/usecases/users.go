package usecases

import (
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
)

// GetUsers usecase returns user list
func GetUsers(userRepo repositories.UserRepository) repositories.Users {
	return userRepo.FetchUsers()
}
