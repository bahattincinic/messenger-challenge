package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Message is a message response
type Message struct {
	gorm.Model

	From    User
	FromID  uint
	To      User
	ToID    uint
	Message string
}

// MessageCreate is a message creation object
type MessageCreate struct {
	Message string `json:"message"`
}

// MessageResponse is a API message response
type MessageResponse struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	FromUser  string    `json:"from"`
	ToUser    string    `json:"to"`
}
