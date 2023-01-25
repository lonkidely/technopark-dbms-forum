package models

import (
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty createthread.go

//easyjson:json
type CreateThreadRequest struct {
	Slug    string    `json:"slug"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Forum   string    `json:"forum"`
}

func NewCreateThreadRequest() *CreateThreadRequest {
	return &CreateThreadRequest{}
}

func (req *CreateThreadRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)
	req.Forum = vars["slug"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *CreateThreadRequest) GetThread() *models.Thread {
	return &models.Thread{
		Slug:    req.Slug,
		Title:   req.Title,
		Author:  req.Author,
		Message: req.Message,
		Created: req.Created,
		Forum:   req.Forum,
	}
}

//easyjson:json
type CreateThreadResponse struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Created time.Time `json:"created"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Message string    `json:"message"`
	Slug    string    `json:"slug"`
	Votes   int       `json:"votes"`
}

func NewCreateThreadResponse(thread *models.Thread) CreateThreadResponse {
	return CreateThreadResponse{
		ID:      thread.ID,
		Title:   thread.Title,
		Author:  thread.Author,
		Forum:   thread.Forum,
		Message: thread.Message,
		Created: thread.Created,
		Votes:   thread.Votes,
		Slug:    thread.Slug,
	}
}
