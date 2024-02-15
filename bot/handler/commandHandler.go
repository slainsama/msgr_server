package handler

import (
	"github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/models"
)

type CommandHandler struct {
	command     string
	handlerFunc func(u *models.TelegramUpdate)
}

func (c *CommandHandler) ShouldHandle(u *models.TelegramUpdate) bool {
	commands := utils.ExtractCommandWithoutArgs(u)
	return len(commands) > 0 && commands[0] == c.command
}

func (c *CommandHandler) HandlerFunc(u *models.TelegramUpdate) {
	c.handlerFunc(u)
}

func NewCommandHandler(command string, handlerFunc func(u *models.TelegramUpdate)) *CommandHandler {
	return &CommandHandler{
		command:     command,
		handlerFunc: handlerFunc,
	}
}
