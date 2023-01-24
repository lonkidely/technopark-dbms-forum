package errors

import (
	"net/http"

	stdErrors "github.com/pkg/errors"
)

var (
	ErrUserExist        = stdErrors.New("such user exist")
	ErrUserNotExist     = stdErrors.New("user doesn't exist")
	ErrEmailAlreadyUsed = stdErrors.New("email already used")
	ErrForumExist       = stdErrors.New("such forum exist")
	ErrForumNotExist    = stdErrors.New("such forum not exist")
	ErrThreadExist      = stdErrors.New("such thread exist")
	ErrThreadNotExist   = stdErrors.New("such thread not exist")
)

func NewErrorClassifier() map[string]int {
	res := make(map[string]int)

	res[ErrUserExist.Error()] = http.StatusConflict
	res[ErrUserNotExist.Error()] = http.StatusNotFound
	res[ErrEmailAlreadyUsed.Error()] = http.StatusConflict
	res[ErrForumExist.Error()] = http.StatusConflict
	res[ErrForumNotExist.Error()] = http.StatusNotFound
	res[ErrThreadExist.Error()] = http.StatusConflict
	res[ErrThreadNotExist.Error()] = http.StatusNotFound

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
