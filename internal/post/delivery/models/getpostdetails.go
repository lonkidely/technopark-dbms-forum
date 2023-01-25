package models

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty getpostdetails.go

type GetPostDetailsRequest struct {
	ID      int
	Related []string
}

func NewGetPostDetailsRequest() *GetPostDetailsRequest {
	return &GetPostDetailsRequest{}
}

func (req *GetPostDetailsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	param := vars["id"]

	req.ID, _ = strconv.Atoi(param)

	param = r.URL.Query().Get("related")

	req.Related = strings.Split(param, ",")

	return nil
}

func (req *GetPostDetailsRequest) GetPost() *models.Post {
	return &models.Post{
		ID: req.ID,
	}
}

func (req *GetPostDetailsRequest) GetParams() *params.PostDetailsParams {
	return &params.PostDetailsParams{
		Related: req.Related,
	}
}

//easyjson:json
type GetPostDetailsAuthorResponse struct {
	Nickname string `json:"nickname,omitempty"`
	FullName string `json:"fullname,omitempty"`
	About    string `json:"about,omitempty"`
	Email    string `json:"email,omitempty"`
}

//easyjson:json
type GetPostDetailsPostResponse struct {
	ID       int       `json:"id,omitempty"`
	Parent   int       `json:"parent,omitempty"`
	Author   string    `json:"author,omitempty"`
	Message  string    `json:"message,omitempty"`
	IsEdited bool      `json:"isEdited,omitempty"`
	Forum    string    `json:"forum,omitempty"`
	Thread   int       `json:"thread,omitempty"`
	Created  time.Time `json:"created,omitempty"`
}

//easyjson:json
type GetPostDetailsThreadResponse struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Slug    string    `json:"slug"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Votes   int       `json:"votes"`
}

//easyjson:json
type GetPostDetailsForumResponse struct {
	Title   string `json:"title"`
	User    string `json:"user"`
	Slug    string `json:"slug"`
	Posts   int    `json:"posts"`
	Threads int    `json:"threads"`
}

//easyjson:json
type PostGetDetailsResponse struct {
	Post   *GetPostDetailsPostResponse   `json:"post"`
	Thread *GetPostDetailsThreadResponse `json:"thread"`
	Author *GetPostDetailsAuthorResponse `json:"author"`
	Forum  *GetPostDetailsForumResponse  `json:"forum"`
}

func NewPostDetailsResponse(postDetails *models.PostDetails) *PostGetDetailsResponse {
	res := &PostGetDetailsResponse{}

	if postDetails.Post.ID != 0 {
		post := GetPostDetailsPostResponse{
			ID:       postDetails.Post.ID,
			Parent:   postDetails.Post.Parent,
			Author:   postDetails.Post.Author,
			Forum:    postDetails.Post.Forum,
			Thread:   postDetails.Post.Thread,
			Message:  postDetails.Post.Message,
			Created:  postDetails.Post.Created,
			IsEdited: postDetails.Post.IsEdited,
		}

		res.Post = &post
	}

	if postDetails.Author.Nickname != "" {
		author := GetPostDetailsAuthorResponse{
			Nickname: postDetails.Author.Nickname,
			FullName: postDetails.Author.FullName,
			About:    postDetails.Author.About,
			Email:    postDetails.Author.Email,
		}

		res.Author = &author
	}

	if postDetails.Thread.ID != 0 {
		thread := GetPostDetailsThreadResponse{
			ID:      postDetails.Thread.ID,
			Title:   postDetails.Thread.Title,
			Author:  postDetails.Thread.Author,
			Forum:   postDetails.Thread.Forum,
			Slug:    postDetails.Thread.Slug,
			Message: postDetails.Thread.Message,
			Created: postDetails.Thread.Created,
			Votes:   postDetails.Thread.Votes,
		}

		res.Thread = &thread
	}

	if postDetails.Forum.User != "" {
		forum := GetPostDetailsForumResponse{
			Title:   postDetails.Forum.Title,
			User:    postDetails.Forum.User,
			Slug:    postDetails.Forum.Slug,
			Posts:   postDetails.Forum.Posts,
			Threads: postDetails.Forum.Threads,
		}

		res.Forum = &forum
	}

	return res
}
