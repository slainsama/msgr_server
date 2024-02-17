package utils

import (
	"strings"

	"github.com/slainsama/msgr_server/models"
)

func ExtractCommands(u *models.TelegramUpdate) ([]string, map[string][]string) {
	// Extract command entities from the update
	var commandEntities []models.Entity
	for _, entity := range u.Message.Entities {
		if entity.Type == "bot_command" {
			commandEntities = append(commandEntities, entity)
		}
	}

	// Extract command and args from the update
	var s, e int
	var command, argsText string
	var commands []string
	args := make(map[string][]string)
	text := u.Message.Text
	for i, commandEntity := range commandEntities {
		s = commandEntity.Offset
		e = s + commandEntity.Length - 1

		command = text[s : e+1]
		commands = append(commands, command)

		if (e + 2) <= len(text)-1 {
			if i < len(commandEntities)-1 {
				argsText = text[e+2 : commandEntities[i+1].Offset-1]
				args[command] = append(args[command], strings.Split(argsText, " ")...)
			} else {
				argsText = text[e+2:]
				args[command] = append(args[command], strings.Split(argsText, " ")...)
			}
		}
	}

	return commands, args
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
