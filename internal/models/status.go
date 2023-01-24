package models

//go:generate easyjson -all -disallow_unknown_fields status.go

type Status struct {
	User   int `json:"user"`
	Forum  int `json:"forum"`
	Thread int `json:"thread"`
	Post   int `json:"post"`
}
