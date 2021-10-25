package router

import (
	ws_handler "go_sample/app/interface/controller/ws"

	"github.com/labstack/echo/v4"
)

func AddWsRoomsRoutingGroup(e *echo.Echo, handler *ws_handler.WsHandler) {
	ws_rooms := e.Group("ws/room")
	{
		ws_rooms.GET("/1", func(c echo.Context) error { return handler.Room().Do(c) })
		ws_rooms.GET("/2", func(c echo.Context) error { return handler.Room().Do(c) })
	}
}
