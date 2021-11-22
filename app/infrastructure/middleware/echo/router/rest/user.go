package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddUsersRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	users := e.Group("users")
	{
		users.GET("", func(c echo.Context) error { return ctrl.User().FindList(c.(*context.CustomContext)) })
		users.GET("/all", func(c echo.Context) error { return ctrl.User().FindAll(c.(*context.CustomContext)) })
		users.GET("/:id", func(c echo.Context) error { return ctrl.User().Find(c.(*context.CustomContext)) })
		users.POST("", func(c echo.Context) error { return ctrl.User().Create(c.(*context.CustomContext)) })
		users.PUT("/:id", func(c echo.Context) error { return ctrl.User().Update(c.(*context.CustomContext)) })
		users.DELETE("/:id", func(c echo.Context) error { return ctrl.User().Delete(c.(*context.CustomContext)) })
	}
}
