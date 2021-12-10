package router

import (
	"fmt"
	ws_controller "go_sample/app/interface/controller/ws"
	enum_connection "go_sample/app/interface/controller/ws/enum/connection"

	"github.com/labstack/echo/v4"
)

func AddWsRoutingGroup(e *echo.Echo, ctrl *ws_controller.WsControllerr) {
	ws := e.Group("ws")
	{
		ws.GET(
			fmt.Sprintf("/%s", string(enum_connection.Chat)),
			func(c echo.Context) error { return ctrl.Chat().Handle(c) },
		)
		ws.GET(
			fmt.Sprintf("/%s", string(enum_connection.DirectMessage)),
			func(c echo.Context) error { return ctrl.DirectMessage().Handle(c) },
		)
		ws.GET(
			fmt.Sprintf("/%s", string(enum_connection.Field)),
			func(c echo.Context) error { return ctrl.Field().Handle(c) },
		)
	}
}
