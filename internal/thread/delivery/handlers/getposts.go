package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type getThreadPostsHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewGetThreadPostsHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &getThreadPostsHandler{
		tu,
	}
}

func (h *getThreadPostsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/thread/{slug_or_id}/posts", h.Action).Methods(http.MethodGet)
}

func (h *getThreadPostsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetThreadPostsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	posts, err := h.threadUsecase.GetPosts(request.GetThread(), request.GetParams())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewThreadGetPostsResponse(posts)

	wrapper.Response(w, http.StatusOK, response)
}
