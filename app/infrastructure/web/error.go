package web

import (
	util_error "go_sample/app/utility/error"
	"net/http"
)

type errorMessage struct {
	Message string
}

type ApiError struct {
	HttpStatusCode int
	Error          errorMessage
}

func NewApiError(err error) *ApiError {
	var httpStatusCode int
	switch err.(type) {
	case *util_error.ErrRecordNotFound:
		httpStatusCode = http.StatusNotFound
	case *util_error.ErrRecordDuplicate:
		httpStatusCode = http.StatusConflict
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	return &ApiError{
		HttpStatusCode: httpStatusCode,
		Error:          errorMessage{Message: err.Error()},
	}
}
