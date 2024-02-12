package botController

import (
	"errors"

	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
	"gorm.io/gorm"
)

// "/start"
func startController(newHandleUpdate models.HandleUpdate) {
	userInfo := newHandleUpdate.NewUpdate.Message.From
	var user models.User
	var message models.Message
	message.ChatId = newHandleUpdate.NewUpdate.Message.Chat.ID
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
			}
			globals.DB.Create(&newUser)
			message.Data = utils.EscapeChar("welcome.")
			botUtils.SendTextMessage(message.ChatId, message.Data)
		}
	} else {
		message.Data = utils.EscapeChar("user already exist.")
		botUtils.SendTextMessage(message.ChatId, message.Data)
	}
}
