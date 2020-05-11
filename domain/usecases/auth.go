package usecases

import (
	"crypto/rand"
	"fmt"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
)

func tokenGenerate() string {
	b := make([]byte, 15)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// CreateAccessToken usecase returns access token
func CreateAccessToken(
	login models.Login,
	authRepo repositories.IAuthRepository,
) (token models.Accesstoken, err error) {

	user, err := authRepo.GetUser(
		login.Username, login.Password,
	)

	if err == nil {
		accessToken := tokenGenerate()
		token = authRepo.CreateAccessToken(accessToken, user)
	}
	return
}

// CreateUser Creates User
func CreateUser(
	signup models.Signup,
	authRepo repositories.IAuthRepository,
) (user models.User) {
	user = authRepo.CreateUser(
		signup.Username, signup.Password, signup.FullName,
	)
	return
}

// CheckAccessToken usecases checks access token
func CheckAccessToken(
	token string, authRepo repositories.IAuthRepository,
	userRepo repositories.IUserRepository,
) (user models.User, err error) {

	accessToken, err := authRepo.CheckAccessToken(token)
	if err != nil {
		return
	}

	user, err = userRepo.FetchUserByID(accessToken.UserID)

	return
}
