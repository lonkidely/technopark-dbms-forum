package models

import "time"

type Post struct {
	ID       int
	Author   string
	Created  time.Time
	Forum    string
	Message  string
	Parent   int
	Thread   int
	IsEdited bool
}
