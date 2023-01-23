package models

//go:generate easyjson -all -disallow_unknown_fields user.go

type User struct {
	ID       int    `json:"-"`
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}
