package botController

import (
	"errors"
	"fmt"
	"github.com/bydBoys/ProcZygoteSDK/config"
	"github.com/slainsama/msgr_server/botUtils"
	"github.com/slainsama/msgr_server/globals"
	"github.com/slainsama/msgr_server/models"
	"github.com/slainsama/msgr_server/scriptUtils"
	"github.com/slainsama/msgr_server/utils"
	"gorm.io/gorm"
	"log"
	"strings"
)

func getUserTaskController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	var message string
	for _, task := range globals.TaskList {
		if task.UserId == userId {
			message = message + task.Id + " " + task.ScriptName
		}
	}
	botUtils.SendTextMessage(userId, message)
}

// "/createTask {scriptName}"
func createUserTaskController(newHandleUpdate models.HandleUpdate) {
	userId := newHandleUpdate.NewUpdate.Message.Chat.ID
	args := newHandleUpdate.Args
	script := new(models.Script)
	if result := globals.DB.First(script, "name = ?", args[0]); result.Error != nil {
		if errors.Is(gorm.ErrRecordNotFound, result.Error) {
			// 记录不存在
			botUtils.SendTextMessage(userId, fmt.Sprintf("no script named %s", args[0]))
			return
		} else {
			// 其他错误
			log.Println(result.Error)
			panic(result.Error)
			return
		}
	}
	if len(args) != len(script.ParamRequired)+1 {
		botUtils.SendTextMessage(userId, fmt.Sprintf("%d params expected,expect:", len(script.ParamRequired)))
		if len(script.ParamRequired) == 0 {
			botUtils.SendTextMessage(userId, "none")
		} else {
			botUtils.SendTextMessage(userId, strings.Join(script.ParamRequired, " "))
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
