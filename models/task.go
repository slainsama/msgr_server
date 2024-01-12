package models

type Task struct {
	Id           string
	UserId       string
	ScriptName   string
	Params       []string
	CallbackData []interface{}
}
