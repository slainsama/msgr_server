package handler

import (
	"regexp"

	"github.com/slainsama/msgr_server/bot/types"
)

// FilterFunc is used to check if this update should be processed by handler.
type FilterFunc func(u *types.TelegramUpdate) bool

var commandRegex = regexp.MustCompile("^/([0-9a-zA-Z_]+)(@[0-9a-zA-Z_]{3,})?")

// Any tells handler to process all updates.
func Any() FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		return true
	}
}

// IsMessage filters updates that look like message (text, photo, location etc.)
func IsMessage() FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		return u.Message != nil
	}
}

// HasPhoto filters updates that contain a photo.
func HasDocument() FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		return u.Message != nil && u.Message.Document != nil
	}
}

// IsAnyCommandMessage filters updates that contain a message and look like a command,
// i. e. have some text and start with a slash ("/").
// If command contains bot username, it is also checked.
func IsAnyCommandMessage() FilterFunc {
	return And(IsMessage(), func(u *types.TelegramUpdate) bool {
		matches := commandRegex.FindStringSubmatch(u.Message.Text)
		return len(matches) != 0
	})
}

// IsCommandMessage filters updates that contain a specific command.
// For example, IsCommandMessage("start") will handle a "/start" command.
// This will also allow the user to pass arguments, e. g. "/start foo bar".
// Commands in format "/start@bot_name" and "/start@bot_name foo bar" are also supported.
// If command contains bot username, it is also checked.
func IsCommandMessage(cmd string) FilterFunc {
	return And(IsAnyCommandMessage(), func(u *types.TelegramUpdate) bool {
		matches := commandRegex.FindStringSubmatch(u.Message.Text)
		actualCmd := matches[1]
		return actualCmd == cmd[1:]
	})
}

// And filters updates that pass ALL of the provided filters.
func And(filters ...FilterFunc) FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		for _, filter := range filters {
			if !filter(u) {
				return false
			}
		}
		return true
	}
}

// Or filters updates that pass ANY of the provided filters.
func Or(filters ...FilterFunc) FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		for _, filter := range filters {
			if filter(u) {
				return true
			}
		}
		return false
	}
}

// Not filters updates that do not pass the provided filter.
func Not(filter FilterFunc) FilterFunc {
	return func(u *types.TelegramUpdate) bool {
		return !filter(u)
	}
}
