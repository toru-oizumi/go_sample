package router

import (
	"fmt"
	ws_handler "go_sample/app/interface/controller/ws"
	enum_connection "go_sample/app/interface/controller/ws/enum/connection"

	"github.com/labstack/echo/v4"
)

func AddWsRoutingGroup(e *echo.Echo, handler *ws_handler.WsHandler) {
	ws := e.Group("ws")
	{
		ws.GET(
			fmt.Sprintf("/%s/:id", string(enum_connection.Chat)),
			func(c echo.Context) error { return handler.Chat().Handle(c) },
		)
		ws.GET(
			fmt.Sprintf("/%s/:id", string(enum_connection.Field)),
			func(c echo.Context) error { return handler.Field().Handle(c) },
		)
	}
}
