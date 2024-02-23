package handler

import (
	"github.com/slainsama/msgr_server/bot/types"
	"github.com/slainsama/msgr_server/bot/utils"
)

type CommandHandler struct {
	command     string
	handlerFunc func(u *types.TelegramUpdate)
}

func (c *CommandHandler) ShouldHandle(u *types.TelegramUpdate) bool {
	commands := utils.ExtractCommandWithoutArgs(u)
	return len(commands) > 0 && commands[0] == c.command
}

func (c *CommandHandler) HandlerFunc(u *types.TelegramUpdate) {
	c.handlerFunc(u)
}

func NewCommandHandler(command string, handlerFunc func(u *types.TelegramUpdate)) *CommandHandler {
	return &CommandHandler{
		command:     command,
		handlerFunc: handlerFunc,
	}
}
