package usecases

import (
	"testing"
	"time"

	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MessagesUsecaseSuite struct {
	suite.Suite
}

func TestMessagesUsecaseInit(t *testing.T) {
	suite.Run(t, new(MessagesUsecaseSuite))
}

func (s *UsersUsecaseSuite) TestCreateMessage() {
	mockedMessageRepo := new(repositories.MockMessageRepository)
	mockedUserRepo := new(repositories.MockUserRepository)
	var fromUser = models.User{
		Model:    gorm.Model{ID: 33},
		Username: "fromuser",
		FullName: "From User",
		Password: "1321323",
	}
	var toUser = models.User{
		Model:    gorm.Model{ID: 35},
		Username: "touser",
		FullName: "To User",
		Password: "te4343st",
	}
	var messageCtx = models.MessageCreate{
		Message: "Hello World",
	}
	var message = models.Message{
		Model:   gorm.Model{ID: 44},
		FromID:  fromUser.ID,
		ToID:    toUser.ID,
		Message: messageCtx.Message,
	}

	mockedUserRepo.On("FetchUserByUsername", toUser.Username).Return(toUser, nil)
	mockedMessageRepo.On("CreateMessage", fromUser, toUser, messageCtx).Return(message)

	createdMessage, err := CreateMessage(
		fromUser, toUser.Username, messageCtx,
		mockedUserRepo, mockedMessageRepo,
	)

	mockedUserRepo.AssertExpectations(s.T())
	mockedMessageRepo.AssertExpectations(s.T())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), createdMessage.ID, message.ID)
	assert.Equal(s.T(), createdMessage.Message, message.Message)
	assert.Equal(s.T(), createdMessage.FromID, message.FromID)
	assert.Equal(s.T(), createdMessage.ToID, message.ToID)
}

func (s *UsersUsecaseSuite) TestGetUserMessages() {
	mockedMessageRepo := new(repositories.MockMessageRepository)
	mockedUserRepo := new(repositories.MockUserRepository)

	var fromUser = models.User{
		Model:    gorm.Model{ID: 33},
		Username: "fromuser",
		FullName: "From User",
		Password: "1321323",
	}
	var toUser = models.User{
		Model:    gorm.Model{ID: 35},
		Username: "touser",
		FullName: "To User",
		Password: "te4343st",
	}
	var message = models.MessageResponse{
		Message:   "hello",
		CreatedAt: time.Now(),
		FromUser:  fromUser.Username,
		ToUser:    toUser.Username,
	}

	var messages repositories.Messages
	messages = append(messages, message)

	mockedUserRepo.On("FetchUserByUsername", toUser.Username).Return(toUser, nil)
	mockedMessageRepo.On("FetchUsersMessages", fromUser, toUser).Return(messages)

	fetchedMessages, err := GetUserMessages(
		fromUser, toUser.Username,
		mockedUserRepo, mockedMessageRepo,
	)

	mockedUserRepo.AssertExpectations(s.T())
	mockedMessageRepo.AssertExpectations(s.T())

	assert.Nil(s.T(), err)
	assert.Equal(s.T(), len(fetchedMessages), len(messages))
	assert.Equal(s.T(), fetchedMessages[0].FromUser, message.FromUser)
	assert.Equal(s.T(), fetchedMessages[0].ToUser, message.ToUser)
	assert.Equal(s.T(), fetchedMessages[0].Message, message.Message)
}
