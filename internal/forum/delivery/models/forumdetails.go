package models

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty forumdetails.go

type ForumDetailsRequest struct {
	Slug string
}

func NewForumDetailsRequest() *ForumDetailsRequest {
	return &ForumDetailsRequest{}
}

func (req *ForumDetailsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.Slug = vars["slug"]

	return nil
}

func (req *ForumDetailsRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

//easyjson:json
type ForumDetailsResponse struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int    `json:"posts"`
	Threads int    `json:"threads"`
}

func NewForumDetailsResponse(forum models.Forum) ForumDetailsResponse {
	return ForumDetailsResponse{
		Title:   forum.Title,
		Slug:    forum.Slug,
		User:    forum.User,
		Posts:   forum.Posts,
		Threads: forum.Threads,
	}
}
