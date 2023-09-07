package wrapper

import (
	"errors"
	"net/http"
)

// ErrInternalServer error
var ErrInternalServer = errors.New("internal server error")

// ErrNotFound error
var ErrNotFound = errors.New("not found")

// ErrBadRequest error
var ErrBadRequest = errors.New("bad request")

// ErrMethodNotAllowed error
var ErrMethodNotAllowed = errors.New("method not allowed")

// GetHTTPStatusCodeByError returns http status code by given error.
//
// Example of return:
//
// When the given error is "errors.ErrNotFound" then the return will be constant value of "http.StatusNotFound".
// Return value will be "http.StatusInternalServerError" when the given error is not included in helper/errors
func getHTTPStatusCodeByError(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInternalServer:
		return http.StatusInternalServerError
	case ErrBadRequest:
		return http.StatusBadRequest
	case ErrMethodNotAllowed:
		return http.StatusMethodNotAllowed
	default:
		return http.StatusInternalServerError
	}
}
