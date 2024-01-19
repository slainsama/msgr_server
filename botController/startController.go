package botController

import (
	"errors"
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"gorm.io/gorm"
)

func startController(newUpdate models.TelegramUpdate) {
	userInfo := newUpdate.Message.From
	var user models.User
	var message models.Message
	message.ChatId = newUpdate.Message.Chat.ID
	result := globals.DB.Where(models.User{ID: userInfo.ID}).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果记录不存在，则创建新记录
			newUser := models.User{
				ID:           userInfo.ID,
				IsBot:        userInfo.IsBot,
				FirstName:    userInfo.FirstName,
				LastName:     userInfo.LastName,
				Username:     userInfo.Username,
				LanguageCode: userInfo.LanguageCode,
				IsAdmin:      false,
				Config:       nil,
			}
			globals.DB.Create(&newUser)
			message.Data = "welcome."
			botUtils.SendTextMessage(message)
		}
	} else {
		message.Data = "user already exist."
		botUtils.SendTextMessage(message)
	}
}
