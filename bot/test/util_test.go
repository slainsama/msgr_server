package test

import (
	"testing"

	"github.com/slainsama/msgr_server/bot/utils"
)

func TestExtractCommands(t *testing.T) {
	update := newHelloUpdate()
	commands, args := utils.ExtractCommands(&update)
	t.Log(commands, args)
}

func TestExtractCommandsWithMultiCommands(t *testing.T) {
	update := newMultiCommandHelloUpdate()
	commands, args := utils.ExtractCommands(&update)
	t.Log(commands, args)
}
