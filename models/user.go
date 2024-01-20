package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           int
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsAdmin      bool
}
