package models

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"

	"lonkidely/technopark-dbms-forum/internal/models"
)

//go:generate easyjson -disallow_unknown_fields -omit_empty createuser.go

//easyjson:json
type CreateUserRequest struct {
	Nickname string
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{}
}

func (req *CreateUserRequest) Bind(r *http.Request) error {
	vars := mux.Vars(r)

	req.Nickname = vars["nickname"]

	body, _ := io.ReadAll(r.Body)

	_ = easyjson.Unmarshal(body, req)

	return nil
}

func (req *CreateUserRequest) GetUser() *models.User {
	return &models.User{
		Nickname: req.Nickname,
		Email:    req.Email,
		About:    req.About,
		FullName: req.FullName,
	}
}

//easyjson:json
type CreateUserResponse struct {
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
	About    string `json:"about"`
	Email    string `json:"email"`
}

//easyjson:json
type CreateUsersResponse []CreateUserResponse

func NewCreateUserResponse(user *models.User) CreateUserResponse {
	return CreateUserResponse{
		Nickname: user.Nickname,
		FullName: user.FullName,
		About:    user.About,
		Email:    user.Email,
	}
}

func NewCreateUsersResponse(users []models.User) CreateUsersResponse {
	response := make([]CreateUserResponse, len(users))

	for idx, val := range users {
		response[idx] = CreateUserResponse{
			Nickname: val.Nickname,
			FullName: val.FullName,
			About:    val.About,
			Email:    val.Email,
		}
	}

	return response
}
