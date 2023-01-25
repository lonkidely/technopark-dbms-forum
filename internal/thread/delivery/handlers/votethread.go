package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
	"lonkidely/technopark-dbms-forum/internal/thread/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/thread/usecase"
)

type voteThreadHandler struct {
	threadUsecase usecase.ThreadUsecase
}

func NewVoteThreadHandler(tu usecase.ThreadUsecase) handler.Handler {
	return &voteThreadHandler{
		threadUsecase: tu,
	}
}

func (h *voteThreadHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/thread/{slug_or_id}/vote", h.Action).Methods(http.MethodPost)
}

func (h *voteThreadHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewVoteThreadRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	thread, err := h.threadUsecase.VoteThread(request.GetThread(), request.GetParams())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewVoteResponse(&thread)

	wrapper.Response(w, http.StatusOK, response)
}
