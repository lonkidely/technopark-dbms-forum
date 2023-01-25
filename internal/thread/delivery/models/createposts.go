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

//go:generate easyjson -disallow_unknown_fields -omit_empty createposts.go

//easyjson:json
type PostRequest struct {
	Parent  int    `json:"parent"`
	Author  string `json:"author"`
	Message string `json:"message"`
}

//easyjson:json
type PostsRequest []PostRequest

type CreatePostsRequest struct {
	SlugOrID string
	Posts    PostsRequest
}

func NewCreatePostsRequest() *CreatePostsRequest {
	return &CreatePostsRequest{}
}

func (req *CreatePostsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, &req.Posts)

	return nil
}

func (req *CreatePostsRequest) GetThread() *models.Thread {
	result := models.Thread{}

	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		result.Slug = req.SlugOrID
	} else {
		result.ID = id
	}

	return &result
}

func (req *CreatePostsRequest) GetPosts() []*models.Post {
	res := make([]*models.Post, len(req.Posts))

	for idx, value := range req.Posts {
		res[idx] = &models.Post{
			Parent:  value.Parent,
			Message: value.Message,
			Author:  value.Author,
		}
	}

	return res
}

//easyjson:json
type PostResponse struct {
	ID       int       `json:"id"`
	Parent   int       `json:"parent"`
	Author   string    `json:"author"`
	Message  string    `json:"message"`
	IsEdited bool      `json:"isEdited"`
	Forum    string    `json:"forum"`
	Thread   int       `json:"thread"`
	Created  time.Time `json:"created"`
}

//easyjson:json
type PostsResponse []PostResponse

func NewCreatePostsResponse(posts []models.Post) PostsResponse {
	res := make([]PostResponse, len(posts))

	for idx, value := range posts {
		res[idx] = PostResponse{
			ID:       value.ID,
			Parent:   value.Parent,
			Author:   value.Author,
			Forum:    value.Forum,
			Thread:   value.Thread,
			Message:  value.Message,
			Created:  value.Created,
			IsEdited: value.IsEdited,
		}
	}

	return res
}
