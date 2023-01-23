package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler interface {
	Action(w http.ResponseWriter, r *http.Request)
	Configure(r *mux.Router)
}
