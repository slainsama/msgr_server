package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int64
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsAdmin      bool
}
