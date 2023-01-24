package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/forum/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/forum/usecase"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
)

type createForumHandler struct {
	forumUsecase usecase.ForumUsecase
}

func NewCreateForumHandler(fu usecase.ForumUsecase) handler.Handler {
	return &createForumHandler{
		forumUsecase: fu,
	}
}

func (h *createForumHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/forum/create", h.Action).Methods(http.MethodPost)
}

func (h *createForumHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewCreateForumRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	forum, err := h.forumUsecase.CreateForum(request.GetForum())
	response := models.NewCreateForumResponse(forum)

	if err != nil {
		if stdErrors.Is(err, errors.ErrForumExist) {
			wrapper.Response(w, http.StatusConflict, response)
			return
		}
		wrapper.ErrorResponse(w, err)
		return
	}

	wrapper.Response(w, http.StatusCreated, response)
}
