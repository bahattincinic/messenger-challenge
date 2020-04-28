package usecases

import (
	"github.com/bahattincinic/messenger-challenge/repositories"
)

// GetUsers usecase returns user list
func GetUsers() repositories.Users {
	return repositories.FetchUsers()
}
