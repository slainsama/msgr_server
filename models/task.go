package models

import "mime/multipart"

type Task struct {
	Id           string
	ZygoteId     string
	UserId       int
	ScriptName   string
	IsLoop       bool
	IsLongTerm   bool
	Params       []string
	CallbackData chan Callback
}

type Callback struct {
	Action string         `json:"action"`
	Msg    string         `json:"msg"`
	File   multipart.File `form:"file"`
}
