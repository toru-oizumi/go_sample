package context

import (
	"encoding/json"
	"go_sample/app/infrastructure/web"
	"go_sample/app/interface/gateway/logger"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"

	util_error "go_sample/app/utility/error"
)

const (
	sessionName      string = "session"
	sessionValueName string = "userID"
	sessionPath      string = "/"
	sessionMaxAge    int    = 86400 * 7
)

type CustomContext struct {
	echo.Context
	RestApiLogger logger.RestApiLogger
}

func (c *CustomContext) CreateErrorResponse(err error) error {
	apiErr := web.NewApiError(err)

	if apiErr.StatusCode != http.StatusInternalServerError {
		// クライアント側起因のエラーはWarningでログを残しておく
		c.RestApiLogger.Warning(apiErr.StatusCode, apiErr.Message)
	} else {
		// InternalServerErrorの場合は、クライアント側でなくサーバー側に問題があるので、エラーログを残す
		c.RestApiLogger.Error(apiErr.StatusCode, apiErr.Message)
	}

	c.JSON(apiErr.StatusCode, apiErr.Message)
	return err
}

func (c *CustomContext) CreateSuccessResponse(statusCode int, entity interface{}) error {
	message, _ := json.Marshal(entity)
	// 正常動作時はDebugレベルで十分
	c.RestApiLogger.Debug(statusCode, string(message))
	c.JSON(statusCode, entity)

	return nil
}

func (c *CustomContext) CheckSession() error {
	if s, err := session.Get(sessionName, c); err != nil {
		return err
	} else {
		if s.Values[sessionValueName] == nil {
			return util_error.NewErrUnauthorized()
		}
		return nil
	}
}

func (c *CustomContext) CreateSession(userID string) error {
	s, _ := session.Get(sessionName, c)
	s.Options = &sessions.Options{
		Path:     sessionPath,
		MaxAge:   sessionMaxAge,
		HttpOnly: true,
		// Secure:   true,
	}
	s.Values[sessionValueName] = userID
	return s.Save(c.Request(), c.Response())
}

func (c *CustomContext) UpdateSession() error {
	s, _ := session.Get(sessionName, c)
	s.Options = &sessions.Options{
		Path:     sessionPath,
		MaxAge:   sessionMaxAge,
		HttpOnly: true,
		// Secure:   true,
	}
	return s.Save(c.Request(), c.Response())
}

func (c *CustomContext) ExpireSession() error {
	s, _ := session.Get(sessionName, c)
	s.Options = &sessions.Options{
		Path:     sessionPath,
		MaxAge:   -1,
		HttpOnly: true,
		// Secure:   true,
	}
	return s.Save(c.Request(), c.Response())
}
