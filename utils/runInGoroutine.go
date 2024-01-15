package utils

import "github.com/slainsama/msgr_server/models"

type WorkerFunction func(newUpdate models.TelegramUpdate)

// RunInGoroutine 接受一个函数类型作为参数，并在新协程中运行该函数
func RunInGoroutine(newUpdate models.TelegramUpdate, worker WorkerFunction) {
	go worker(newUpdate)
}
