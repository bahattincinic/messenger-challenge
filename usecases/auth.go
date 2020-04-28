package usecases

import (
	"github.com/bahattincinic/messenger-challenge/models"
	"github.com/bahattincinic/messenger-challenge/repositories"
)

// CreateAccessToken usecase returns access token
func CreateAccessToken(login models.Login) (token models.Accesstoken, err error) {
	accessToken, err := repositories.CreateAccessToken(
		login.Username, login.Password,
	)

	if err == nil {
		token = models.Accesstoken{
			Token: accessToken,
		}
	}
	return
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
