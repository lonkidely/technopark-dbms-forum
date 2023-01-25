package models

import "time"

//go:generate easyjson -all -disallow_unknown_fields thread.go

type Thread struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Slug    string    `json:"slug"`
	Votes   int       `json:"votes"`
}
