package models

import "gorm.io/gorm"

type configData map[string]string

type userConfig map[string]configData

type User struct {
	gorm.Model
	ID           int
	IsBot        bool
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsAdmin      bool
	Config       userConfig
}
