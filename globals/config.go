package globals

import (
	"github.com/slainsama/msgr_server/models"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var UnmarshaledConfig models.Config

func initConfig() {
	configFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal(configFile, &UnmarshaledConfig)
	if err != nil {
		log.Println(err)
	}
	log.Println(Config)
}
