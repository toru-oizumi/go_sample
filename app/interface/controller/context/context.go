package context

import (
	"go_sample/app/interface/gateway/logger"
	"net/http"
)

type Context interface {
	Request() *http.Request
	Param(string) string
	QueryParam(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	Validate(interface{}) error
	CreateErrorResponse(logger.RestApiLogger, error) error
	CreateSuccessResponse(logger.RestApiLogger, int, interface{}) error
}
