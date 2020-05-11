package repositories

import (
	"github.com/bahattincinic/messenger-challenge/domain/models"
	"github.com/jinzhu/gorm"
)

// MessageRepository ..
type MessageRepository struct {
	db *gorm.DB
}

// Messages represents list of messages
type Messages []models.MessageResponse

// IMessageRepository is a interface of MessageRepository
type IMessageRepository interface {
	CreateMessage(
		fromUser models.User, toUser models.User, messageContext models.MessageCreate,
	) (message models.Message)

	FetchUsersMessages(
		fromUser models.User, toUser models.User,
	) Messages
}

// NewMessageRepo ..
func NewMessageRepo(db *gorm.DB) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

// CreateMessage inserts message to the database
func (r *MessageRepository) CreateMessage(
	fromUser models.User, toUser models.User,
	messageContext models.MessageCreate,
) (message models.Message) {

	message = models.Message{
		FromID:  fromUser.ID,
		ToID:    toUser.ID,
		Message: messageContext.Message,
	}
	r.db.Create(&message)

	message.From = fromUser
	message.To = toUser

	return
}

// FetchUsersMessages returns messages
func (r *MessageRepository) FetchUsersMessages(
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
