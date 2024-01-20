package init

import (
	"github.com/slainsama/msgr_server/botController"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/scriptController"
)

func Init() {
	globals.Init()
	botController.Init()
	scriptController.Init()
}
