package scriptUtils

import (
	"github.com/google/uuid"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
)

func TaskCreate(Task models.Task) (newTask models.Task) {
	newId := uuid.New().String()
	Task.Id = newId
	globals.TaskList[newId] = newTask
	return newTask
}
