package models

// User Model Object
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`
}

// Login Object
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Signup Object
type Signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"fullname"`
}
