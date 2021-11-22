package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddPlaysRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	plays := e.Group("plays")
	{
		plays.GET("", func(c echo.Context) error { return ctrl.Play().FindAll(c.(*context.CustomContext)) })
		plays.GET("/:id", func(c echo.Context) error { return ctrl.Play().Find(c.(*context.CustomContext)) })
		plays.POST("", func(c echo.Context) error { return ctrl.Play().Create(c.(*context.CustomContext)) })
		plays.PUT("/:id", func(c echo.Context) error { return ctrl.Play().Update(c.(*context.CustomContext)) })
		plays.DELETE("/:id", func(c echo.Context) error { return ctrl.Play().Delete(c.(*context.CustomContext)) })
	}
}
