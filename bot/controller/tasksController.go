package controller

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/bydBoys/ProcZygoteSDK/config"
	"github.com/slainsama/msgr_server/bot/botMethod"
	models2 "github.com/slainsama/msgr_server/bot/models"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/scriptUtils"
	"github.com/slainsama/msgr_server/utils"
	"gorm.io/gorm"
)

func GetUserTaskController(newHandleUpdate models2.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	var message string
	for _, task := range globals.TaskList {
		if task.UserId == userId {
			message = message + task.Id + " " + task.ScriptName
		}
	}
	botMethod.SendTextMessage(userId, message)
}

// "/createTask {scriptName}"
func CreateUserTaskController(newHandleUpdate models2.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	args := newHandleUpdate.Args
	script := new(models.Script)
	if result := globals.DB.First(script, "name = ?", args[0]); result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			// 记录不存在
			botMethod.SendTextMessage(userId, fmt.Sprintf("no script named %s", args[0]))
			return
		} else {
			// 其他错误
			log.Println(result.Error)
			panic(result.Error)
			return
		}
	}
	if len(args) != len(script.ParamRequired)+1 {
		botMethod.SendTextMessage(userId, fmt.Sprintf("%d params expected,expect:", len(script.ParamRequired)))
		if len(script.ParamRequired) == 0 {
			botMethod.SendTextMessage(userId, "none")
		} else {
			botMethod.SendTextMessage(userId, strings.Join(script.ParamRequired, " "))
		}
		return
	}
	newTask := scriptUtils.TaskCreate(userId)
	scriptCommand, _ := utils.FormatCommand(script.Command, newTask.Id, newTask.ScriptName, args[1:])
	zygoteUuid, _ := globals.Zygote.StartProcess(scriptCommand, config.UserIsolated{Enable: false}, config.CGroup{
		Enable:      false,
		CpuShare:    "",
		CpuSet:      "",
		MemoryLimit: "",
	})
	newTask.ZygoteId = zygoteUuid
}
