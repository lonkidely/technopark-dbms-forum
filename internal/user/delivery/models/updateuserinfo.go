package models

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty updateuserinfo.go

//easyjson:json
type UpdateUserInfoRequest struct {
	Nickname string
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

func NewUpdateUserInfoRequest() *UpdateUserInfoRequest {
	return &UpdateUserInfoRequest{}
}

func (req *UpdateUserInfoRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.Nickname = vars["nickname"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *UpdateUserInfoRequest) GetUser() *models.User {
	return &models.User{
		Nickname: req.Nickname,
		Email:    req.Email,
		About:    req.About,
		FullName: req.FullName,
	}
}

//easyjson:json
type UpdateUserInfoResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

func NewUpdateUserInfoResponse(user *models.User) UpdateUserInfoResponse {
	return UpdateUserInfoResponse{
		Nickname: user.Nickname,
		FullName: user.FullName,
		About:    user.About,
		Email:    user.Email,
	}
}
