package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddAuthenticationRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	authentications := e.Group("")
	{
		authentications.POST("/sign-in", func(c echo.Context) error { return ctrl.Authentication().SingIn(c.(*context.CustomContext)) })
		authentications.POST("/sign-up", func(c echo.Context) error { return ctrl.Authentication().SingUp(c.(*context.CustomContext)) })
		authentications.POST("/activate", func(c echo.Context) error { return ctrl.Authentication().Activate(c.(*context.CustomContext)) })
		authentications.POST("/change-password", func(c echo.Context) error { return ctrl.Authentication().ChangePassword(c.(*context.CustomContext)) })
		authentications.DELETE("/sign-out", func(c echo.Context) error { return ctrl.Authentication().SingOut(c.(*context.CustomContext)) })
	}
}
