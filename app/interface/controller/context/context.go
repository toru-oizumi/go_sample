package context

import (
	"net/http"
)

type Context interface {
	Request() *http.Request
	Param(string) string
	QueryParam(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	Validate(interface{}) error
	CreateErrorResponse(error) error
	CreateSuccessResponse(int, interface{}) error
}
