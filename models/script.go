package models

type Script struct {
	id            string
	name          string
	status        string
	paramRequired []string
	dataReturn    []string
}
