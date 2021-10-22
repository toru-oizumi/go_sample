package router

import (
	"go_sample/app/interface/controller"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddGroupsRoutingGroup(e *echo.Echo, ctrl *controller.Controller) {
	groups := e.Group("groups")
	{
		groups.GET("", func(c echo.Context) error { return ctrl.Group().FindAll(c.(*context.CustomContext)) })
		groups.GET("/:id", func(c echo.Context) error { return ctrl.Group().Find(c.(*context.CustomContext)) })
		groups.POST("", func(c echo.Context) error { return ctrl.Group().Create(c.(*context.CustomContext)) })
		groups.PUT("/:id", func(c echo.Context) error { return ctrl.Group().Update(c.(*context.CustomContext)) })
		groups.DELETE("/:id", func(c echo.Context) error { return ctrl.Group().Delete(c.(*context.CustomContext)) })
	}
}
