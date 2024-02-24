package controller

import (
	"errors"

	"github.com/slainsama/msgr_server/bot/botMethod"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/handler"
	"github.com/slainsama/msgr_server/bot/types"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"gorm.io/gorm"
)

func init() {
	botGlobals.Dispatcher.AddHandler(handler.NewCommandHandler("/start", startController))
}

// startController "/start"
func startController(u *types.TelegramUpdate) int {
	userInfo := u.Message.From

	var user models.User
	result := globals.DB.Where(models.User{ID: userInfo.ID}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// 如果记录不存在，则创建新记录
			newUser := models.User{
				ID:           userInfo.ID,
				IsBot:        userInfo.IsBot,
				FirstName:    userInfo.FirstName,
				LastName:     userInfo.LastName,
				Username:     userInfo.UserName,
				LanguageCode: userInfo.LanguageCode,
				IsAdmin:      false,
			}
			globals.DB.Create(&newUser)
			botMethod.SendTextMessage(u.Message.From.ID, botUtils.EscapeChar("welcome."))
		}
	} else {
		botMethod.SendTextMessage(u.Message.From.ID, botUtils.EscapeChar("user already exist."))
	}
	return handler.HandleSuccess
}
