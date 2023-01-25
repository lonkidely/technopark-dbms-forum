package models

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty getposts.go

type GetThreadPostsRequest struct {
	SlugOrID string
	Limit    int
	Since    int
	Desc     bool
	Sort     string
}

func NewGetThreadPostsRequest() *GetThreadPostsRequest {
	return &GetThreadPostsRequest{}
}

func (req *GetThreadPostsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	value := r.URL.Query().Get("limit")
	if value == "" {
		req.Limit = 100
	} else {
		req.Limit, _ = strconv.Atoi(value)
	}

	value = r.URL.Query().Get("since")
	if value == "" {
		req.Since = -1
	} else {
		req.Since, _ = strconv.Atoi(value)
	}

	var err error
	value = r.URL.Query().Get("desc")
	if req.Desc, err = strconv.ParseBool(value); err != nil {
		req.Desc = false
	}

	value = r.URL.Query().Get("sort")
	if value == "" {
		req.Sort = "flat"
	} else {
		req.Sort = value
	}

	return nil
}

func (req *GetThreadPostsRequest) GetThread() *models.Thread {
	result := models.Thread{}

	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		result.Slug = req.SlugOrID
	} else {
		result.ID = id
	}

	return &result
}

func (req *GetThreadPostsRequest) GetParams() *params.GetPostsParams {
	return &params.GetPostsParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
		Sort:  req.Sort,
	}
}

//easyjson:json
type GetThreadPostResponse struct {
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
type GetThreadPostsResponse []GetThreadPostResponse

func NewThreadGetPostsResponse(posts []models.Post) GetThreadPostsResponse {
	res := make([]GetThreadPostResponse, len(posts))

	for idx, value := range posts {
		res[idx] = GetThreadPostResponse{
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
