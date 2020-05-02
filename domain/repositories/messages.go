package repositories

import (
	"time"

	"github.com/bahattincinic/messenger-challenge/domain/models"
)

// CreateMessage inserts message to the database
func CreateMessage(fromUser models.User, toUser models.User, message models.MessageCreate) models.Message {
	now := time.Now()

	messageID := InsertRow(
		"INSERT INTO Messages (from_user, to_user, message, created_at) VALUES (?,?,?,?)",
		fromUser.ID, toUser.ID, message.Message, now,
	)

	return models.Message{
		ID:        messageID,
		From:      fromUser.Username,
		To:        toUser.Username,
		Message:   message.Message,
		CreatedAt: now.String(),
	}
}

// Messages represents list of messages
type Messages []models.Message

// FetchUsersMessages returns messages
func FetchUsersMessages(fromUser models.User, toUser models.User) Messages {
	var messages Messages
	rows := fetchRows(
		`SELECT m.id, f_user.username, t_user.username, m.message, m.created_at
		FROM Messages as m
		INNER JOIN Users as f_user ON (f_user.id = m.from_user)
		INNER JOIN Users as t_user ON (t_user.id = m.to_user)
		WHERE (from_user = ? AND to_user = ?) OR (from_user = ? AND to_user = ?)`,
		fromUser.ID, toUser.ID, toUser.ID, fromUser.ID,
	)

	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.From, &message.To,
			&message.Message, &message.CreatedAt)
		CheckErr(err)
		messages = append(messages, message)
	}

	return messages
}
