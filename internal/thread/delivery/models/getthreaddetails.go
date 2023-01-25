package models

import (
	"lonkidely/technopark-dbms-forum/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty getthreaddetails.go

//easyjson:json
type GetThreadDetailsRequest struct {
	SlugOrID string
}

func NewGetThreadDetailsRequest() *GetThreadDetailsRequest {
	return &GetThreadDetailsRequest{}
}

func (req *GetThreadDetailsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	return nil
}

func (req *GetThreadDetailsRequest) GetThread() *models.Thread {
	result := models.Thread{}

	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		result.Slug = req.SlugOrID
	} else {
		result.ID = id
	}

	return &result
}

//easyjson:json
type GetThreadDetailsResponse struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Slug    string    `json:"slug"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Votes   int       `json:"votes"`
}

func NewGetThreadDetailsResponse(thread *models.Thread) *GetThreadDetailsResponse {
	return &GetThreadDetailsResponse{
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
