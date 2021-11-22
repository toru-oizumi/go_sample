package context

import (
	"encoding/json"
	"go_sample/app/infrastructure/web"
	"go_sample/app/interface/gateway/logger"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) CreateErrorResponse(logger logger.RestApiLogger, err error) error {
	apiErr := web.NewApiError(err)

	if apiErr.StatusCode != http.StatusInternalServerError {
		// クライアント側起因のエラーはWarningでログを残しておく
		logger.Warning(apiErr.StatusCode, apiErr.Message)
	} else {
		// InternalServerErrorの場合は、クライアント側でなくサーバー側に問題があるので、エラーログを残す
		logger.Error(apiErr.StatusCode, apiErr.Message)
	}

	c.JSON(apiErr.StatusCode, apiErr.Message)
	return err
}

func (c *CustomContext) CreateSuccessResponse(logger logger.RestApiLogger, status_code int, entity interface{}) error {
	message, _ := json.Marshal(entity)
	// 正常動作時はDebugレベルで十分
	logger.Debug(status_code, string(message))
	c.JSON(status_code, entity)
	return nil
}
