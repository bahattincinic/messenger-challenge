package models

import "github.com/jinzhu/gorm"

// Accesstoken Model Object
type Accesstoken struct {
	gorm.Model

	ID     uint `gorm:"primary_key";json:"-"`
	Token  string
	User   User
	UserID uint
}
