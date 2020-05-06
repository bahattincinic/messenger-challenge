package usecases

import (
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
)

// CreateMessage inserts message to Database
func CreateMessage(
	fromUser models.User, toUser string,
	message models.MessageCreate,
	userRepo repositories.UserRepository,
	messageRepo repositories.MessageRepisotry,
) (createdMessage models.Message, err error) {

	user, err := userRepo.FetchUserByUsername(toUser)
	if err != nil {
		return
	}

	createdMessage = messageRepo.CreateMessage(fromUser, user, message)
	return
}

// GetUserMessages returns users messages
func GetUserMessages(
	fromUser models.User,
	toUser string,
	userRepo repositories.UserRepository,
	messageRepo repositories.MessageRepisotry,
) (messages repositories.Messages, err error) {

	user, err := userRepo.FetchUserByUsername(toUser)
	if err != nil {
		return
	}

	messages = messageRepo.FetchUsersMessages(fromUser, user)
	return
}
