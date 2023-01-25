package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/post/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/post/usecase"
)

type updatePostHandler struct {
	postUsecase usecase.PostUsecase
}

func NewUpdatePostHandler(pu usecase.PostUsecase) handler.Handler {
	return &updatePostHandler{
		pu,
	}
}

func (h *updatePostHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/post/{id}/details", h.Action).Methods(http.MethodPost)
}

func (h *updatePostHandler) Action(w http.ResponseWriter, r *http.Request) {
	req := models.NewUpdatePostRequest()

	bindError := req.Bind(r)
	if bindError != nil {
		return
	}

	resultPost, err := h.postUsecase.UpdatePost(req.GetPost())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewUpdatePostResponse(resultPost)

	wrapper.Response(w, http.StatusOK, response)
}
