package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type getThreadDetailsHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewGetThreadDetailsHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &getThreadDetailsHandler{
		threadUsecase: tu,
	}
}

func (h *getThreadDetailsHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/thread/{slug_or_id}/details", h.Action).Methods(http.MethodGet)
}

func (h *getThreadDetailsHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetThreadDetailsRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	thread, err := h.threadUsecase.GetThreadDetails(request.GetThread())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewGetThreadDetailsResponse(&thread)

	wrapper.Response(w, http.StatusOK, response)
}
