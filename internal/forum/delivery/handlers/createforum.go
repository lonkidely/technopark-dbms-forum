package handlers

import (
	"lonkidely/technopark-dbms-forum/internal/forum/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/forum/usecase"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
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
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewCreateForumResponse(forum)

	wrapper.Response(w, http.StatusCreated, response)
}
