package main

import (
	"github.com/fvbock/endless"
	"github.com/slainsama/msgr_server/inits"
	"github.com/slainsama/msgr_server/server"
	"log"
)

func main() {
	inits.Init()
	endlessServer := endless.NewServer("0.0.0.0:8081", server.Server)
	err := endlessServer.ListenAndServe()
	if err != nil {
		log.Println("something wrong with starting.")
	}

}
