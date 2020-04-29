package usecases

import (
	"crypto/rand"
	"fmt"

	"github.com/bahattincinic/messenger-challenge/models"
	"github.com/bahattincinic/messenger-challenge/repositories"
)

func tokenGenerate() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// CreateAccessToken usecase returns access token
func CreateAccessToken(login models.Login) (token models.Accesstoken, err error) {
	user, err := repositories.GetUser(
		login.Username, login.Password,
	)

	if err == nil {
		accessToken := tokenGenerate()
		token = repositories.CreateAccessToken(accessToken, user)
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

// CheckAccessToken usecases checks access token
func CheckAccessToken(token string) (user models.User, err error) {
	accessToken, err := repositories.CheckAccessToken(token)

	if err != nil {
		return
	}

	user, err = repositories.FetchUserByID(accessToken.UserID)

	return
}
