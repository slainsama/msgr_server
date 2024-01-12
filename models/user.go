package models

import "gorm.io/gorm"

type configData map[string]string

type userConfig map[string]configData

type User struct {
	gorm.Model
	id             string
	TelegramChatId int
	IsAdmin        bool
	Config         userConfig
}
