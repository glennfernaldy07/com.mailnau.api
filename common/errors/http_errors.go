package errors

import "net/http"

type Error struct {
	Code    int
	Message string
	Cause   error
}

func newError(code int, msg string, cause error) *Error {
	return &Error{code, msg, cause}
}

func (e *Error) Error() string {
	return e.Cause.Error()
}

func NewBadRequestError(msg string, cause error) *Error {
	return newError(http.StatusBadRequest, msg, cause)
}

func NewInternalError(cause error) *Error {
	return newError(http.StatusInternalServerError, "Kesalahan internal server", cause)
}
