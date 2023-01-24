package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/forum/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/forum/usecase"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
)

type getForumThreadsHandler struct {
	forumUsecase usecase.ForumUsecase
}

func NewGetForumThreadsHandler(fu usecase.ForumUsecase) handler.Handler {
	return &getForumThreadsHandler{
		forumUsecase: fu,
	}
}

func (h *getForumThreadsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/forum/{slug}/threads", h.Action).Methods(http.MethodGet)
}

func (h *getForumThreadsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetForumThreadsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	forum, err := h.forumUsecase.GetForumThreads(request.GetForum(), request.GetParams())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewGetForumThreadsResponse(forum)

	wrapper.Response(w, http.StatusOK, response)
}
