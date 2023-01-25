package models

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty updatethreaddetails.go

//easyjson:json
type UpdateThreadDetailsRequest struct {
	SlugOrID string
	Title    string `json:"title"`
	Message  string `json:"message"`
}

func NewUpdateThreadDetailsRequest() *UpdateThreadDetailsRequest {
	return &UpdateThreadDetailsRequest{}
}

func (req *UpdateThreadDetailsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *UpdateThreadDetailsRequest) GetThread() *models.Thread {
	result := models.Thread{}

	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		result.Slug = req.SlugOrID
	} else {
		result.ID = id
	}

	result.Title = req.Title
	result.Message = req.Message

	return &result
}

//easyjson:json
type UpdateThreadDetailsResponse struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Slug    string    `json:"slug"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Votes   int       `json:"votes"`
}

func NewUpdateThreadDetailsResponse(thread *models.Thread) *UpdateThreadDetailsResponse {
	return &UpdateThreadDetailsResponse{
		ID:      thread.ID,
		Title:   thread.Title,
		Author:  thread.Author,
		Forum:   thread.Forum,
		Slug:    thread.Slug,
		Message: thread.Message,
		Created: thread.Created,
		Votes:   thread.Votes,
	}
}
