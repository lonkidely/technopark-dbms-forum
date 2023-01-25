package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type createPostsHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewCreatePostsHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &createPostsHandler{
		tu,
	}
}

func (h *createPostsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/thread/{slug_or_id}/create", h.Action).Methods(http.MethodPost)
}

func (h *createPostsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewCreatePostsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	posts, err := h.threadUsecase.CreatePosts(request.GetThread(), request.GetPosts())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewCreatePostsResponse(posts)

	wrapper.Response(w, http.StatusCreated, response)
}
