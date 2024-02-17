package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/bydBoys/ProcZygoteSDK/config"
	"github.com/slainsama/msgr_server/bot/botMethod"
	botGlobals "github.com/slainsama/msgr_server/bot/globals"
	"github.com/slainsama/msgr_server/bot/handler"
	botUtils "github.com/slainsama/msgr_server/bot/utils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/utils"
	"gorm.io/gorm"
)

func init() {
	botGlobals.Dispatcher.AddHandler(handler.NewCommandHandler("/createTask", createUserTaskController))
	botGlobals.Dispatcher.AddHandler(handler.NewCommandHandler("/tasks", getUserTaskController))
}

func getUserTaskController(u *models.TelegramUpdate) {
	userId := u.Message.Chat.ID
	var message string
	for _, task := range globals.TaskList {
		if task.UserId == userId {
			message = message + task.Id + " " + task.ScriptName
		}
	}
	botMethod.SendTextMessage(userId, message)
}

// createUserTaskController "/createTask {scriptName}"
func createUserTaskController(u *models.TelegramUpdate) {
	userId := u.Message.Chat.ID

	commands, messageArgs := botUtils.ExtractCommands(u)
	args := messageArgs[commands[0]]
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
		}
	}
	/*
		argsNum := utils.GetArgsNum(script)
		if len(args) != argsNum+1 {
			botMethod.SendTextMessage(userId, fmt.Sprintf("%d params expected,expect:", argsNum))
			if argsNum == 0 {
				botMethod.SendTextMessage(userId, "none")
			} else {
				botMethod.SendTextMessage(userId, strings.Join(argsNum, " "))
			}
			return
		}
	*/
	newTask := utils.TaskCreate(userId)
	scriptCommand, _ := utils.FormatCommand(script.Command, newTask.Id, newTask.ScriptName, args[1:])
	zygoteUuid, _ := globals.Zygote.StartProcess(scriptCommand, config.UserIsolated{Enable: false}, config.CGroup{
		Enable:      false,
		CpuShare:    "",
		CpuSet:      "",
		MemoryLimit: "",
	})
	newTask.ZygoteId = zygoteUuid
}
