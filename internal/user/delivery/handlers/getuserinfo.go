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

type getUserInfoHandler struct {
	userUsecase usecase.UserUsecase
}

func NewGetUserInfoHandler(u usecase.UserUsecase) handler.Handler {
	return &getUserInfoHandler{
		userUsecase: u,
	}
}

func (h *getUserInfoHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/user/{nickname}/profile", h.Action).Methods(http.MethodGet)
}

func (h *getUserInfoHandler) Action(w http.ResponseWriter, r *http.Request) {
	req := models.NewGetUserInfoRequest()

	bindError := req.Bind(r)
	if bindError != nil {
		return
	}

	requestUser := req.GetUser()

	user, err := h.userUsecase.GetUserInfo(requestUser)
	if err != nil {
		if stdErrors.Is(err, errors.ErrUserNotExist) {
			wrapper.ErrorResponse(w, err)
			return
		}
		return
	}

	response := models.NewGetUserInfoResponse(&user)

	wrapper.Response(w, http.StatusOK, response)
}
