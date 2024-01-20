package main

import (
	"github.com/fvbock/endless"
	"github.com/slainsama/msgr_server/init"
	"github.com/slainsama/msgr_server/server"
	"log"
)

func main() {
	init.Init()
	endlessServer := endless.NewServer("0.0.0.0:8081", server.Server)
	err := endlessServer.ListenAndServe()
	if err != nil {
		log.Println("something wrong with starting.")
	}

}
