package main

import (
	"github.com/fvbock/endless"
	"github.com/slainsama/msgr_server/botController"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/scriptController"
	"github.com/slainsama/msgr_server/server"
	"log"
)

func init() {
	globals.Init()
	server.Init()
	botController.Init()
	scriptController.Init()
}

func main() {
	endlessServer := endless.NewServer("0.0.0.0:8081", server.Server)
	err := endlessServer.ListenAndServe()
	if err != nil {
		log.Println("something wrong with starting.")
	}

}
