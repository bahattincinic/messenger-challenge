package repositories

import (
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/stretchr/testify/mock"
)

// MockAuthRepository is a mock implementation for unit tests
type MockAuthRepository struct {
	mock.Mock
}

// CreateAccessToken is a mocked func
func (m *MockAuthRepository) CreateAccessToken(accessToken string, user models.User) models.Accesstoken {
	args := m.Called(accessToken, user)

	return args.Get(0).(models.Accesstoken)
}

// GetUser is a mocked func
func (m *MockAuthRepository) GetUser(username string, password string) (user models.User, err error) {
	args := m.Called(username, password)

	return args.Get(0).(models.User), args.Error(1)
}

// CreateUser is a mocked func
func (m *MockAuthRepository) CreateUser(username string, password string, fullname string) models.User {
	args := m.Called(username, password, fullname)

	return args.Get(0).(models.User)
}

// CheckAccessToken is a mocked func
func (m *MockAuthRepository) CheckAccessToken(token string) (accessToken models.Accesstoken, err error) {
	args := m.Called(token)

	return args.Get(0).(models.Accesstoken), args.Error(1)
}

// MockUserRepository is a mock implementation for unit tests
type MockUserRepository struct {
	mock.Mock
}

// FetchUsers is a mocked func
func (m *MockUserRepository) FetchUsers() Users {
	args := m.Called()

	return args.Get(0).(Users)
}

// FetchUserByID is a mocked func
func (m *MockUserRepository) FetchUserByID(userID uint) (user models.User, err error) {
	args := m.Called(userID)

	return args.Get(0).(models.User), args.Error(1)
}

// FetchUserByUsername is a mocked func
func (m *MockUserRepository) FetchUserByUsername(username string) (user models.User, err error) {
	args := m.Called(username)

	return args.Get(0).(models.User), args.Error(1)
}

// MockMessageRepository is a mock implementation for unit tests
type MockMessageRepository struct {
	mock.Mock
}

// CreateMessage is a mocked func
func (m *MockMessageRepository) CreateMessage(
	fromUser models.User, toUser models.User,
	messageContext models.MessageCreate,
) (message models.Message) {
	args := m.Called(fromUser, toUser, messageContext)

	return args.Get(0).(models.Message)
}

// FetchUsersMessages is a mocked func
func (m *MockMessageRepository) FetchUsersMessages(
	fromUser models.User, toUser models.User,
) Messages {
	args := m.Called(fromUser, toUser)

	return args.Get(0).(Messages)
}
