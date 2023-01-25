package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type updateThreadDetailsHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewUpdateThreadDetailsHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &updateThreadDetailsHandler{
		threadUsecase: tu,
	}
}

func (h *updateThreadDetailsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/thread/{slug_or_id}/details", h.Action).Methods(http.MethodPost)
}

func (h *updateThreadDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewUpdateThreadDetailsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	thread, err := h.threadUsecase.UpdateThreadDetails(request.GetThread())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewUpdateThreadDetailsResponse(&thread)

	wrapper.Response(w, http.StatusOK, response)
}
