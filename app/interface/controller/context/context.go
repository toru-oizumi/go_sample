package context

import (
	"go_sample/app/interface/controller/logger"
)

type Context interface {
	Param(string) string
	Bind(interface{}) error
	JSON(int, interface{}) error
	CreateErrorResponse(logger.Logger, error) error
	CreateSuccessResponse(logger.Logger, int, interface{}) error
}
