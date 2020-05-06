package models

import "github.com/jinzhu/gorm"

// User Model Object
type User struct {
	gorm.Model

	Username string `json:"username"`
	FullName string `json:"fullname"`
	Password string `json:"-"`
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
