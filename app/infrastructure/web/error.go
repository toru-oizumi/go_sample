package web

import (
	util_error "go_sample/app/utility/error"
	"net/http"
)

type ApiError struct {
	StatusCode            int     `json:"statusCode"`
	ErrorTitle            string  `json:"errorTitle"`
	Detail                string  `json:"detail"`
	MessagesToBeDisplayed *string `json:"messagesToBeDisplayed"` // 画面上で表示したい文言を指定する場合に使用する
}

func NewApiError(err error) ApiError {
	var httpStatusCode int
	var errorTitle string
	var messagesToBeDisplayed *string

	switch e := err.(type) {
	case util_error.ErrValidationError:
		httpStatusCode = http.StatusBadRequest
	case util_error.ErrActivationNotRequired:
		httpStatusCode = http.StatusBadRequest
		errorTitle = e.Title

	case util_error.ErrAuthenticationFailed:
		httpStatusCode = http.StatusUnauthorized
		errorTitle = e.Title
	case util_error.ErrActivationRequired:
		httpStatusCode = http.StatusUnauthorized
		errorTitle = e.Title

	case util_error.ErrChangePasswordRequired:
		httpStatusCode = http.StatusUnauthorized
		errorTitle = e.Title

	case util_error.ErrEmailAlreadyExists:
		httpStatusCode = http.StatusConflict
		errorTitle = e.Title

	case util_error.ErrUnauthorizedRequest:
		httpStatusCode = http.StatusUnauthorized
		errorTitle = e.Title
	case util_error.ErrEntityNotExists:
		httpStatusCode = http.StatusNotFound
		errorTitle = e.Title
	case util_error.ErrEntityAlreadyExists:
		httpStatusCode = http.StatusConflict
		errorTitle = e.Title
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	return ApiError{
		StatusCode:            httpStatusCode,
		ErrorTitle:            errorTitle,
		Detail:                err.Error(),
		MessagesToBeDisplayed: messagesToBeDisplayed,
	}
}
