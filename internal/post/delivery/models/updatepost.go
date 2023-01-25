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

//go:generate easyjson -disallow_unknown_fields -omit_empty updatepost.go

//easyjson:json
type UpdatePostRequest struct {
	ID      int
	Message string `json:"message"`
}

func NewUpdatePostRequest() *UpdatePostRequest {
	return &UpdatePostRequest{}
}

func (req *UpdatePostRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	param := vars["id"]

	req.ID, _ = strconv.Atoi(param)

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *UpdatePostRequest) GetPost() *models.Post {
	return &models.Post{
		ID:      req.ID,
		Message: req.Message,
	}
}

//easyjson:json
type UpdatePostResponse struct {
	ID       int       `json:"id"`
	Parent   int       `json:"parent"`
	Author   string    `json:"author"`
	Message  string    `json:"message"`
	IsEdited bool      `json:"isEdited"`
	Forum    string    `json:"forum"`
	Thread   int       `json:"thread"`
	Created  time.Time `json:"created"`
}

func NewUpdatePostResponse(post *models.Post) *UpdatePostResponse {
	return &UpdatePostResponse{
		ID:       post.ID,
		Parent:   post.Parent,
		Author:   post.Author,
		Forum:    post.Forum,
		Thread:   post.Thread,
		Message:  post.Message,
		Created:  post.Created,
		IsEdited: post.IsEdited,
	}
}
