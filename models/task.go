package models

type Task struct {
	Id           string
	UserId       int
	ScriptName   string
	IsLoop       bool
	IsLongTerm   bool
	Params       []string
	CallbackData chan Callback
}

type Callback struct {
	Action string
	Data   interface{}
}
