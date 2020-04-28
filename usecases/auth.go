package usecases

import (
	"github.com/bahattincinic/messenger-challenge/models"
	"github.com/bahattincinic/messenger-challenge/repositories"
)

// CreateAccessToken usecase returns access token
func CreateAccessToken(login models.Login) models.User {
	return models.User{
		Username: "Test",
		ID:       1,
		FullName: "Test",
	}
}

// CreateUser Creates User
func CreateUser(signup models.Signup) models.User {
	userID := repositories.CreateUser(
		signup.Username, signup.Password, signup.FullName,
	)

	return models.User{
		Username: signup.Username,
		ID:       userID,
		FullName: signup.FullName,
	}
}
