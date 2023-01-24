package models

//go:generate easyjson -all -disallow_unknown_fields forum.go

type Forum struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int    `json:"posts"`
	Threads int    `json:"threads"`
}
