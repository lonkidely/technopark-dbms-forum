package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/service/usecase"
)

type statusHandler struct {
	serviceUsecase usecase.ServiceUsecase
}

func NewStatusHandler(serviceUsecase usecase.ServiceUsecase) handler.Handler {
	return &statusHandler{
		serviceUsecase: serviceUsecase,
	}
}

func (h *statusHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/service/status", h.Action).Methods(http.MethodGet)
}

func (h *statusHandler) Action(w http.ResponseWriter, r *http.Request) {
	response, err := h.serviceUsecase.Status()

	if err != nil {
		return
	}

	wrapper.Response(w, http.StatusOK, response)
}
