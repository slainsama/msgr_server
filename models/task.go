package models

type Task struct {
	Id           string
	UserId       string
	ScriptName   string
	IsLoop       bool
	IsLongTerm   bool
	Params       []string
	CallbackData []interface{}
}
