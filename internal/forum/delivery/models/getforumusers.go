package models

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty getforumusers.go

type GetForumUsersRequest struct {
	Slug  string
	Limit int
	Since string
	Desc  bool
}

func NewGetForumUsersRequest() *GetForumUsersRequest {
	return &GetForumUsersRequest{}
}

func (req *GetForumUsersRequest) Bind(r *http.Request) error {
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

func (req *GetForumUsersRequest) GetForum() *models.Forum {
	return &models.Forum{
		Slug: req.Slug,
	}
}

func (req *GetForumUsersRequest) GetParams() *params.GetForumUsersParams {
	return &params.GetForumUsersParams{
		Limit: req.Limit,
		Since: req.Since,
		Desc:  req.Desc,
	}
}

//easyjson:json
type GetForumUserResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type GetForumUsersResponse []GetForumUserResponse

func NewGetForumUsersResponse(users []*models.User) GetForumUsersResponse {
	response := make([]GetForumUserResponse, len(users))

	for idx, val := range users {
		response[idx] = GetForumUserResponse{
			Nickname: val.Nickname,
			FullName: val.FullName,
			About:    val.About,
			Email:    val.Email,
		}
	}

	return response
}
