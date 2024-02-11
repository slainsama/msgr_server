package scriptUtils

import (
	"github.com/google/uuid"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
)

func TaskCreate(userId int) (task models.Task) {
	var newTask models.Task
	newId := uuid.New().String()
	newTask.Id = newId
	newTask.UserId = userId
	globals.TaskList[newId] = newTask
	return newTask
}
