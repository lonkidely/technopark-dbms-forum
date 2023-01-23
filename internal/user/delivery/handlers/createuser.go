package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/user/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/user/usecase"
)

type createUserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewCreateUserHandler(u usecase.UserUsecase) handler.Handler {
	return &createUserHandler{
		userUsecase: u,
	}
}

func (h *createUserHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/user/{nickname}/create", h.Action).Methods(http.MethodPost)
}

func (h *createUserHandler) Action(w http.ResponseWriter, r *http.Request) {
	req := models.NewCreateUserRequest()

	bindError := req.Bind(r)
	if bindError != nil {
		return
	}

	requestUser := req.GetUser()

	users, err := h.userUsecase.CreateUser(requestUser)

	if err != nil {
		if stdErrors.Is(err, errors.ErrUserExist) {
			response := models.NewCreateUsersResponse(users)

			wrapper.Response(w, http.StatusConflict, response)

			return
		}
	}

	response := models.NewCreateUserResponse(&users[0])

	wrapper.Response(w, http.StatusCreated, response)
}
