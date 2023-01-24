package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/service/usecase"
)

type clearHandler struct {
	serviceUsecase usecase.ServiceUsecase
}

func NewClearHandler(serviceUsecase usecase.ServiceUsecase) handler.Handler {
	return &clearHandler{
		serviceUsecase: serviceUsecase,
	}
}

func (h *clearHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/service/clear", h.Action).Methods(http.MethodPost)
}

func (h *clearHandler) Action(w http.ResponseWriter, r *http.Request) {
	err := h.serviceUsecase.Clear()

	if err != nil {
		return
	}

	wrapper.NoBody(w, http.StatusOK)
}
