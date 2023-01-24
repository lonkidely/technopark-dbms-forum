package models

import (
	"io"
	"net/http"

	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty createforum.go

//easyjson:json
type CreateForumRequest struct {
	Title string `json:"title"`
	User  string `json:"user"`
	Slug  string `json:"slug"`
}

func NewCreateForumRequest() *CreateForumRequest {
	return &CreateForumRequest{}
}

func (req *CreateForumRequest) Bind(r *http.Request) error {
	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *CreateForumRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug:  req.Slug,
		Title: req.Title,
		User:  req.User,
	}
}

//easyjson:json
type CreateForumResponse struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int    `json:"posts"`
	Threads int    `json:"threads"`
}

func NewCreateForumResponse(forum models.Forum) CreateForumResponse {
	return CreateForumResponse{
		Title:   forum.Title,
		Slug:    forum.Slug,
		User:    forum.User,
		Posts:   forum.Posts,
		Threads: forum.Threads,
	}
}
