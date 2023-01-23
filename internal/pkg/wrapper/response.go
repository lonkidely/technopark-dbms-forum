package wrapper

import (
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"net/http"

	"github.com/mailru/easyjson"
)

func getEasyJSON(someStruct interface{}) ([]byte, error) {
	someStructUpdate, _ := someStruct.(easyjson.Marshaler)

	out, err := easyjson.Marshal(someStructUpdate)

	return out, err
}

func Response(w http.ResponseWriter, statusCode int, someStruct interface{}) {
	out, err := getEasyJSON(someStruct)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	_, err = w.Write(out)
	if err != nil {
		return
	}
}

func NoBody(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
}

//go:generate easyjson -disallow_unknown_fields -omit_empty response.go

//easyjson:json
type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}

func ErrorResponse(w http.ResponseWriter, err error) {
	errMsg := ErrorMessage{
		Message: err.Error(),
	}
	Response(w, errors.GetErrorCode(err), errMsg)
}
