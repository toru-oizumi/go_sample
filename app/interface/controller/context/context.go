package context

import (
	"go_sample/app/domain/model"
	"net/http"
)

type Context interface {
	Request() *http.Request
	Param(name string) string
	QueryParam(name string) string
	Bind(i interface{}) error
	JSON(code int, i interface{}) error
	Validate(i interface{}) error

	CheckSession() error
	GetUserIDFromSession() (*model.UserID, error)
	CreateSession(userID string) error
	UpdateSession() error
	ExpireSession() error

	CreateErrorResponse(err error) error
	CreateSuccessResponse(statusCode int, entity interface{}) error
}
