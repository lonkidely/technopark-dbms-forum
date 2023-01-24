package models

import (
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty forumdetails.go

type GetForumThreadsRequest struct {
	Slug  string
	Limit int
	Since string
	Desc  bool
}

func NewGetForumThreadsRequest() *GetForumThreadsRequest {
	return &GetForumThreadsRequest{}
}

func (req *GetForumThreadsRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.Slug = vars["slug"]

	value := r.URL.Query().Get("limit")
	if value == "" {
		req.Limit = 100
	} else {
		req.Limit, _ = strconv.Atoi(value)
	}

	req.Since = r.URL.Query().Get("since")

	var err error
	value = r.URL.Query().Get("desc")
	if req.Desc, err = strconv.ParseBool(value); err != nil {
		req.Desc = false
	}

	return nil
}

func (req *GetForumThreadsRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

func (req *GetForumThreadsRequest) GetParams() *params.GetForumThreadsParams {
	return &params.GetForumThreadsParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
	}
}

//easyjson:json
type GetForumThreadResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Forum   string `json:"forum"`
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Created string `json:"created"`
	Votes   int    `json:"votes"`
}

//easyjson:json
type GetForumThreadsResponse []GetForumThreadResponse

func NewGetForumThreadsResponse(threads []*models.Thread) GetForumThreadsResponse {
	response := make([]GetForumThreadResponse, len(threads))

	for idx, val := range threads {
		response[idx] = GetForumThreadResponse{
			ID:      val.ID,
			Title:   val.Title,
			Author:  val.Author,
			Forum:   val.Forum,
			Slug:    val.Slug,
			Message: val.Message,
			Created: val.Created,
			Votes:   val.Votes,
		}
	}

	return response
}
