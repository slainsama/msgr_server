package utils

import (
	"github.com/slainsama/msgr_server/bot/models"
)

type WorkerFunction func(newHandleUpdate models.HandleUpdate)

// RunInGoroutine 接受一个函数类型作为参数，并在新协程中运行该函数
func RunInGoroutine(newHandleUpdate models.HandleUpdate, worker WorkerFunction) {
	go worker(newHandleUpdate)
}
