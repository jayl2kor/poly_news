package polyNewsErrors

import (
	"errors"
	"fmt"
	"net/http"
)

type PolyNewsError struct {
	error
	Code    int
	Message string
}

func New(err error, code int, msg string) PolyNewsError {
	return PolyNewsError{
		error:   err,
		Code:    code,
		Message: msg,
	}
}

func Newf(err error, code int, format string, args ...string) PolyNewsError {
	return New(err, code, fmt.Sprintf(format, args))
}

func Wrap(err error) error {
	var polyNewError PolyNewsError
	if errors.As(err, &polyNewError) {
		return polyNewError
	}
	return PolyNewsError{
		error: err,
		Code:  http.StatusInternalServerError,
	}
}
