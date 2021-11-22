package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddGroupsRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	groups := e.Group("groups")
	{
		groups.GET("", func(c echo.Context) error { return ctrl.Group().FindList(c.(*context.CustomContext)) })
		groups.GET("/all", func(c echo.Context) error { return ctrl.Group().FindAll(c.(*context.CustomContext)) })
		groups.GET("/:id", func(c echo.Context) error { return ctrl.Group().Find(c.(*context.CustomContext)) })
		groups.POST("", func(c echo.Context) error { return ctrl.Group().Create(c.(*context.CustomContext)) })
		groups.PUT("/:id", func(c echo.Context) error { return ctrl.Group().Update(c.(*context.CustomContext)) })
		groups.DELETE("/:id", func(c echo.Context) error { return ctrl.Group().Delete(c.(*context.CustomContext)) })
	}
}
