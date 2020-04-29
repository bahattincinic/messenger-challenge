package models

// Accesstoken Model Object
type Accesstoken struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}
