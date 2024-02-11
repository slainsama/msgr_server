package globals

import (
	"github.com/bydBoys/ProcZygoteSDK/ProZygote"
	"log"
)

var Zygote ProZygote.ProZygote

func initZygote() {
	Zygote := new(ProZygote.ProZygote)
	if err := Zygote.Init("127.0.0.1:9963"); err != nil {
		log.Fatal(err)
		return
	}
	clientVersion, serverVersion, err := Zygote.GetVersion()
	if err != nil {
		log.Fatal("call GetVersion error ", err)
		return
	}
	log.Println("client version: ", clientVersion)
	log.Println("server version: ", serverVersion)
}
