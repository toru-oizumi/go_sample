package web

import (
	util_error "go_sample/app/utility/error"
	"net/http"
)

type ApiError struct {
	StatusCode int
	Message    string
}

func NewApiError(err error) ApiError {
	var httpStatusCode int
	switch err.(type) {
	case util_error.ErrValidationError:
		httpStatusCode = http.StatusBadRequest
	case util_error.ErrUnauthorized:
		httpStatusCode = http.StatusUnauthorized
	case util_error.ErrRecordNotFound:
		httpStatusCode = http.StatusNotFound
	case util_error.ErrRecordDuplicate:
		httpStatusCode = http.StatusConflict
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	return ApiError{
		StatusCode: httpStatusCode,
		Message:    err.Error(),
	}
}
