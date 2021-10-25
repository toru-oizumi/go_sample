package router

import (
	"go_sample/app/interface/controller"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddRoomsRoutingGroup(e *echo.Echo, ctrl *controller.Controller) {
	rooms := e.Group("rooms")
	{
		rooms.GET("", func(c echo.Context) error { return ctrl.Room().FindAll(c.(*context.CustomContext)) })
		rooms.GET("/:id", func(c echo.Context) error { return ctrl.Room().Find(c.(*context.CustomContext)) })
		rooms.POST("", func(c echo.Context) error { return ctrl.Room().Create(c.(*context.CustomContext)) })
		rooms.PUT("/:id", func(c echo.Context) error { return ctrl.Room().Update(c.(*context.CustomContext)) })
		rooms.DELETE("/:id", func(c echo.Context) error { return ctrl.Room().Delete(c.(*context.CustomContext)) })
	}
}
