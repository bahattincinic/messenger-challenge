package usecases

import (
	"testing"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersUsecaseSuite struct {
	suite.Suite
}

func TestUsersUsecaseInit(t *testing.T) {
	suite.Run(t, new(UsersUsecaseSuite))
}

func (s *UsersUsecaseSuite) TestGetUsers() {
	mockedUserRepo := new(repositories.MockUserRepository)
	var users repositories.Users
	var user = models.User{
		Model:    gorm.Model{ID: 33},
		Username: "test",
		FullName: "test",
		Password: "test",
	}
	users = append(users, user)

	mockedUserRepo.On("FetchUsers").Return(users)

	fetchedUsers := GetUsers(mockedUserRepo)
	mockedUserRepo.AssertExpectations(s.T())
	assert.Equal(s.T(), len(fetchedUsers), len(users))
	assert.Equal(s.T(), fetchedUsers[0].Username, user.Username)
	assert.Equal(s.T(), fetchedUsers[0].FullName, user.FullName)
	assert.Equal(s.T(), fetchedUsers[0].ID, user.ID)
}
