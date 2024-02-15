package models

// Script "{scriptName}" to get scriptName;{arg1} to get args
type Script struct {
	Id      string
	Name    string
	Command string //such as "python3 {scriptName} {arg1} {arg2}"
	Status  string
}
