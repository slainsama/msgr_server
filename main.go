package main

import (
	"log"

	"github.com/fvbock/endless"
	"github.com/slainsama/msgr_server/bot"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/scriptController"
	"github.com/slainsama/msgr_server/server"
)

func init() {
	globals.Init()
	server.Init()
	bot.Init()
	scriptController.Init()
}

func main() {
	config := globals.UnmarshaledConfig
	if config.DEBUG.Switch {
		log.SetFlags(log.LstdFlags | log.Llongfile)
	}

	endlessServer := endless.NewServer("0.0.0.0:8081", server.Server)
	err := endlessServer.ListenAndServe()
	if err != nil {
		log.Println("something wrong with starting.")
	}

}
