package models

// Message is a message response
type Message struct {
	ID        int64  `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Message   string `json:"message"`
	CreatedAt string `json:"created_at"`
}

// MessageCreate is a message creation object
type MessageCreate struct {
	Message string `json:"message"`
}
