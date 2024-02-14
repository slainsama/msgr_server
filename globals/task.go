package globals

import (
	"github.com/slainsama/msgr_server/models"
)

var MessageChannel = make(chan models.TelegramUpdate)

var TaskList map[string]models.Task

func initTaskList() {
	TaskList = make(map[string]models.Task)
}
