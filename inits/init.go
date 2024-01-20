package inits

import (
	"github.com/slainsama/msgr_server/botController"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/scriptController"
	"github.com/slainsama/msgr_server/server"
)

func Init() {
	globals.Init()
	server.Init()
	botController.Init()
	scriptController.Init()
}
