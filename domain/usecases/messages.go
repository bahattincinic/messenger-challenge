package usecases

import (
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/bahattincinic/messenger-challenge/domain/repositories"
)

// CreateMessage inserts message to Database
func CreateMessage(fromUser models.User, toUser string, message models.MessageCreate) (createdMessage models.Message, err error) {
	user, err := repositories.FetchUserByUsername(toUser)

	if err != nil {
		return
	}

	createdMessage = repositories.CreateMessage(fromUser, user, message)
	return
}

// GetUserMessages returns users messages
func GetUserMessages(fromUser models.User, toUser string) (messages repositories.Messages, err error) {
	user, err := repositories.FetchUserByUsername(toUser)

	if err != nil {
		return
	}

	messages = repositories.FetchUsersMessages(fromUser, user)
	return
}
