package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	ErrUserExist        = stdErrors.New("such user exist")
	ErrUserNotExist     = stdErrors.New("user doesn't exist")
	ErrEmailAlreadyUsed = stdErrors.New("email already used")
)

func NewErrorClassifier() map[string]int {
	res := make(map[string]int)

	res[ErrUserNotExist.Error()] = http.StatusConflict
	res[ErrUserNotExist.Error()] = http.StatusNotFound
	res[ErrEmailAlreadyUsed.Error()] = http.StatusConflict

	return res
}

var errorClassifier = NewErrorClassifier()

func GetErrorCode(err error) int {
	res, exist := errorClassifier[err.Error()]
	if !exist {
		return http.StatusConflict
	}
	return res
}
