package context

import (
	"go_sample/app/interface/controller/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Context interface {
	Request() *http.Request
	// echoの型に依存しているが、そもそもecho用のContext Interfaceなので許容する
	Response() *echo.Response
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	Validate(interface{}) error
	BindAndValidate(logger.Logger, interface{}) error
	CreateErrorResponse(logger.Logger, error) error
	CreateSuccessResponse(logger.Logger, int, interface{}) error
}
