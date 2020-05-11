package usecases

import (
	"testing"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type AuthUsecaseSuite struct {
	suite.Suite
}

func TestAuthUsecaseInit(t *testing.T) {
	suite.Run(t, new(AuthUsecaseSuite))
}

func (s *AuthUsecaseSuite) TestCreateAccessToken() {
	mockedRepo := new(repositories.MockAuthRepository)

	var login = models.Login{
		Username: "test",
		Password: "123456",
	}
	var user = models.User{
		Model:    gorm.Model{ID: 11},
		Username: login.Username,
		FullName: "User 11",
		Password: login.Password,
	}
	var token = models.Accesstoken{
		Token:  "1121-211",
		User:   user,
		UserID: user.ID,
	}

	mockedRepo.On("GetUser", login.Username, login.Password).Return(user, nil)
	mockedRepo.On("CreateAccessToken", mock.Anything, user).Return(token, nil)

	createdToken, err := CreateAccessToken(login, mockedRepo)

	mockedRepo.AssertExpectations(s.T())
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), token.Token, createdToken.Token)
	assert.Equal(s.T(), token.UserID, createdToken.UserID)
}

func (s *AuthUsecaseSuite) TestCreateUser() {
	mockedRepo := new(repositories.MockAuthRepository)

	var signup = models.Signup{
		Username: "example",
		Password: "1q2w3e4f",
		FullName: "Example Fullname",
	}
	var user = models.User{
		Model:    gorm.Model{ID: 12},
		Username: signup.Username,
		FullName: signup.FullName,
		Password: signup.Password,
	}

	mockedRepo.On("CreateUser", signup.Username, signup.Password, signup.FullName).Return(user)

	createdUser := CreateUser(signup, mockedRepo)

	mockedRepo.AssertExpectations(s.T())
	assert.Equal(s.T(), createdUser.ID, user.ID)
	assert.Equal(s.T(), createdUser.Username, user.Username)
	assert.Equal(s.T(), createdUser.FullName, user.FullName)
}

func (s *AuthUsecaseSuite) TesCheckAccessToken() {
	mockedAuthRepo := new(repositories.MockAuthRepository)
	mockedUserRepo := new(repositories.MockUserRepository)

	var accessToken = models.Accesstoken{
		Model:  gorm.Model{ID: 33},
		Token:  "1212a2121s",
		UserID: 33,
	}
	var user = models.User{
		Model:    gorm.Model{ID: 33},
		Username: "test",
		FullName: "test",
		Password: "test",
	}

	mockedAuthRepo.On("CheckAccessToken", accessToken.Token).Return(accessToken, nil)
	mockedUserRepo.On("FetchUserByID", user.ID).Return(user, nil)

	checkedUser, err := CheckAccessToken(accessToken.Token, mockedAuthRepo, mockedUserRepo)

	mockedAuthRepo.AssertExpectations(s.T())
	mockedUserRepo.AssertExpectations(s.T())
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), checkedUser.ID, user.ID)
	assert.Equal(s.T(), checkedUser.Username, user.Username)
	assert.Equal(s.T(), checkedUser.FullName, user.FullName)
}
