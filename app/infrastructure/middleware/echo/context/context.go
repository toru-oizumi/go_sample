package context

import (
	"encoding/json"
	"go_sample/app/infrastructure/web"
	"go_sample/app/interface/controller/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) CreateErrorResponse(logger logger.Logger, err error) error {
	apiErr := web.NewApiError(err)
	c.JSON(apiErr.HttpStatusCode, apiErr.Error)

	if apiErr.HttpStatusCode != http.StatusInternalServerError {
		// クライアント側起因のエラーはWarningでログを残しておく
		logger.Warning(apiErr.HttpStatusCode, apiErr.Error.Message)
	} else {
		// InternalServerErrorの場合は、クライアント側でなくサーバー側に問題があるので、エラーログを残す
		logger.Error(apiErr.HttpStatusCode, apiErr.Error.Message)
	}

	return err
}

func (c *CustomContext) CreateSuccessResponse(logger logger.Logger, http_status_code int, entity interface{}) error {
	message, _ := json.Marshal(entity)
	// 正常動作時はDebugレベルで十分
	logger.Debug(http_status_code, string(message))
	c.JSON(http_status_code, entity)
	return nil
}
