package utils

import (
	"strings"

	"github.com/slainsama/msgr_server/models"
)

func ExtractCommands(newUpdate models.TelegramUpdate) ([]string, map[string][]string) {
	var commands []models.Entity
	var messages []string
	var startPos int
	var endPos int
	argsPos := make(map[string][2]int)
	args := make(map[string][]string)
	for _, entity := range newUpdate.Message.Entities {
		if entity.Type == "bot_command" {
			commands = append(commands, entity)
		}
	}
	text := newUpdate.Message.Text
	for _, command := range commands {
		startPos = command.Offset
		endPos = startPos + command.Length
		argsPos[text[startPos:endPos]] = [2]int{startPos, endPos}
		messages = append(messages, text[startPos:endPos])
	}
	messageSize := len(messages)
	for messagePos, message := range messages {
		var argsEndPos int
		argsStartPos := argsPos[message][1]
		if messagePos > messageSize {
			argsEndPos = len(newUpdate.Message.Text) - 1
		} else {
			argsEndPos = argsPos[messages[messagePos+1]][0]
		}
		argString := newUpdate.Message.Text[argsStartPos:argsEndPos]
		argsResult := strings.Fields(argString)
		args[message] = argsResult
	}
	return messages, args
}

func ExtractCommandWithoutArgs(u *models.TelegramUpdate) []string {
	var commandEntities []models.Entity
	var commands []string

	for _, entity := range u.Message.Entities {
		if entity.Type == "bot_command" {
			commandEntities = append(commandEntities, entity)
		}
	}

	text := u.Message.Text
	for _, command := range commandEntities {
		startPos := command.Offset
		endPos := startPos + command.Length
		commands = append(commands, text[startPos:endPos])
	}
	return commands
}
