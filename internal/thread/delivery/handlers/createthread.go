package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type createThreadHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewCreateThreadHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &createThreadHandler{
		threadUsecase: tu,
	}
}

func (h *createThreadHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/forum/{slug}/create", h.Action).Methods(http.MethodPost)
}

func (h *createThreadHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewCreateThreadRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	thread, err := h.threadUsecase.CreateThread(request.GetThread())
	response := models.NewCreateThreadResponse(&thread)

	if err != nil {
		if stdErrors.Is(err, errors.ErrThreadExist) {
			wrapper.Response(w, http.StatusConflict, response)
			return
		}
		wrapper.ErrorResponse(w, err)
		return
	}

	wrapper.Response(w, http.StatusCreated, response)
}
