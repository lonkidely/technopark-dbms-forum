package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"lonkidely/technopark-dbms-forum/internal/forum/delivery/models"
	"lonkidely/technopark-dbms-forum/internal/forum/usecase"
	"lonkidely/technopark-dbms-forum/internal/pkg/handler"
	"lonkidely/technopark-dbms-forum/internal/pkg/wrapper"
)

type getForumUsersHandler struct {
	forumUsecase usecase.ForumUsecase
}

func NewGetForumUsersHandler(fu usecase.ForumUsecase) handler.Handler {
	return &getForumUsersHandler{
		forumUsecase: fu,
	}
}

func (h *getForumUsersHandler) Configure(r *mux.Router) {
	r.HandleFunc("/api/forum/{slug}/users", h.Action).Methods(http.MethodGet)
}

func (h *getForumUsersHandler) Action(w http.ResponseWriter, r *http.Request) {
	request := models.NewGetForumUsersRequest()

	errBind := request.Bind(r)
	if errBind != nil {
		return
	}

	forum, err := h.forumUsecase.GetForumUsers(request.GetForum(), request.GetParams())
	if err != nil {
		wrapper.ErrorResponse(w, err)
		return
	}

	response := models.NewGetForumUsersResponse(forum)

	wrapper.Response(w, http.StatusOK, response)
}
