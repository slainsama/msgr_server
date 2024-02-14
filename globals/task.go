package globals

import (
	models2 "github.com/slainsama/msgr_server/bot/models"
	"github.com/slainsama/msgr_server/models"
)

var MessageChannel = make(chan models2.TelegramUpdate)

var TaskList map[string]models.Task

func initTaskList() {
	TaskList = make(map[string]models.Task)
}
