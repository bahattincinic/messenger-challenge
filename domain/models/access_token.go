package models

import "github.com/jinzhu/gorm"

// Accesstoken Model Object
type Accesstoken struct {
	gorm.Model

	Token  string
	User   User
	UserID uint
}
