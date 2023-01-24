package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/forum/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/forum/usecase"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
)

type forumDetailsHandler struct {
	forumUsecase usecase.ForumUsecase
}

func NewForumDetailsHandler(fu usecase.ForumUsecase) handler.Handler {
	return &forumDetailsHandler{
		forumUsecase: fu,
	}
}

func (h *forumDetailsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/forum/{slug}/details", h.Action).Methods(http.MethodGet)
}

func (h *forumDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewForumDetailsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	forum, err := h.forumUsecase.GetForumDetails(request.GetForum())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewForumDetailsResponse(forum)

	wrapper.Response(w, http.StatusOK, response)
}
