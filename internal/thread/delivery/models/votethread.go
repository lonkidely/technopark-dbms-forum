package models

import (
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty votethread.go

//easyjson:json
type VoteThreadRequest struct {
	SlugOrID string
	Nickname string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func NewVoteThreadRequest() *VoteThreadRequest {
	return &VoteThreadRequest{}
}

func (req *VoteThreadRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.SlugOrID = vars["slug_or_id"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *VoteThreadRequest) GetThread() *models.Thread {
	result := models.Thread{}

	id, err := strconv.Atoi(req.SlugOrID)
	if err != nil {
		result.Slug = req.SlugOrID
	} else {
		result.ID = id
	}

	return &result
}

func (req *VoteThreadRequest) GetParams() *params.VoteThreadParams {
	return &params.VoteThreadParams{
		Nickname: req.Nickname,
		Voice:    req.Voice,
	}
}

//easyjson:json
type VoteThreadResponse struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Author  string    `json:"author"`
	Forum   string    `json:"forum"`
	Slug    string    `json:"slug"`
	Message string    `json:"message"`
	Created time.Time `json:"created"`
	Votes   int       `json:"votes"`
}

func NewVoteResponse(thread *models.Thread) *VoteThreadResponse {
	return &VoteThreadResponse{
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
