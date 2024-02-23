package models

import "mime/multipart"

type Task struct {
	Id           string
	ZygoteId     string
	UserId       int64
	ScriptName   string
	IsLoop       bool
	IsLongTerm   bool
	Params       []string
	CallbackData chan Callback
}

type Callback struct {
	Action string         `json:"action"` //"sendText" or "sendPhoto"
	Msg    string         `json:"msg"`
	File   multipart.File `form:"file"`
}
