package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddDirectMessagesRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	chats := e.Group("direct-messages")
	{
		chats.GET("", func(c echo.Context) error { return ctrl.DirectMessage().FindAll(c.(*context.CustomContext)) })
	}
}
