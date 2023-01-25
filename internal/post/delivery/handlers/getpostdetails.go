package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/post/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/post/usecase"
)

type getPostDetailsHandler struct {
	postUsecase usecase.PostUsecase
}

func NewGetPostDetailsHandler(pu usecase.PostUsecase) handler.Handler {
	return &getPostDetailsHandler{
		pu,
	}
}

func (h *getPostDetailsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/post/{id}/details", h.Action).Methods(http.MethodGet)
}

func (h *getPostDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	req := models.NewGetPostDetailsRequest()

	bindError := req.Bind(r)
	if bindError != nil {
		return
	}

	postDetails, err := h.postUsecase.GetPostDetails(req.GetPost(), req.GetParams())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewPostDetailsResponse(postDetails)

	wrapper.Response(w, http.StatusOK, response)
}
