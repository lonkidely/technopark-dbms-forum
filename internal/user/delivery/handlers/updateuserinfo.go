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

type updateUserInfoHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUpdateUserInfoHandler(u usecase.UserUsecase) handler.Handler {
	return &updateUserInfoHandler{
		userUsecase: u,
	}
}

func (h *updateUserInfoHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/user/{nickname}/profile", h.Action).Methods(http.MethodPost)
}

func (h *updateUserInfoHandler) Action(w http.ResponseWriter, r *http.Request) {
	req := models.NewUpdateUserInfoRequest()

	bindError := req.Bind(r)
	if bindError != nil {
		return
	}

	requestUser := req.GetUser()

	user, err := h.userUsecase.UpdateUser(requestUser)
	if err != nil {
		if stdErrors.Is(err, errors.ErrUserNotExist) {
			wrapper.ErrorResponse(w, err)
			return
		}
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewUpdateUserInfoResponse(&user)

	wrapper.Response(w, http.StatusOK, response)
}
