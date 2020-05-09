package repositories

import (
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/jinzhu/gorm"
)

// MessageRepisotry ..
type MessageRepisotry struct {
	db *gorm.DB
}

// NewMessageRepo ..
func NewMessageRepo(db *gorm.DB) *MessageRepisotry {
	return &MessageRepisotry{
		db: db,
	}
}

// CreateMessage inserts message to the database
func (r *MessageRepisotry) CreateMessage(
	fromUser models.User, toUser models.User,
	messageContext models.MessageCreate,
) (message models.Message) {

	message = models.Message{
		FromID:  fromUser.ID,
		From:    fromUser,
		ToID:    toUser.ID,
		To:      toUser,
		Message: messageContext.Message,
	}
	r.db.Create(&message)

	return
}

// Messages represents list of messages
type Messages []models.MessageResponse

// FetchUsersMessages returns messages
func (r *MessageRepisotry) FetchUsersMessages(
	fromUser models.User, toUser models.User,
) Messages {

	var messages Messages

	r.db.Table("messages").Select(
		`messages.message, messages.created_at,
		from_user.username as from_user, to_user.username as to_user`,
	).Where(
		"(from_id = ? AND to_id = ?)",
		fromUser.ID, toUser.ID,
	).Or(
		"(from_id = ? AND to_id = ?)",
		toUser.ID, fromUser.ID,
	).Joins(
		"JOIN users as from_user ON from_user.id = messages.from_id",
	).Joins(
		"JOIN users as to_user ON to_user.id = messages.to_id",
	).Scan(&messages)

	return messages
}
