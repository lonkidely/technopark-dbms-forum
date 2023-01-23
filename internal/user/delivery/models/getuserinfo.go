package models

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty getuserinfo.go

type GetUserInfoRequest struct {
	Nickname string
}

func NewGetUserInfoRequest() *GetUserInfoRequest {
	return &GetUserInfoRequest{}
}

func (req *GetUserInfoRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.Nickname = vars["nickname"]

	return nil
}

func (req *GetUserInfoRequest) GetUser() *models.User {
	return &models.User{
		Nickname: req.Nickname,
	}
}

//easyjson:json
type GetUserInfoResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

func NewGetUserInfoResponse(user *models.User) GetUserInfoResponse {
	return GetUserInfoResponse{
		Nickname: user.Nickname,
		FullName: user.FullName,
		About:    user.About,
		Email:    user.Email,
	}
}
