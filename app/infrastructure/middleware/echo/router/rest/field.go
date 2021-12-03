package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddFieldsRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	fields := e.Group("fields")
	{
		fields.GET("", func(c echo.Context) error { return ctrl.Field().FindAll(c.(*context.CustomContext)) })
		fields.GET("/:id", func(c echo.Context) error { return ctrl.Field().Find(c.(*context.CustomContext)) })
		fields.POST("", func(c echo.Context) error { return ctrl.Field().Create(c.(*context.CustomContext)) })
		fields.PUT("/:id", func(c echo.Context) error { return ctrl.Field().Update(c.(*context.CustomContext)) })
		fields.DELETE("/:id", func(c echo.Context) error { return ctrl.Field().Delete(c.(*context.CustomContext)) })
	}
}
