package router

import (
	rest_controller "go_sample/app/interface/controller/rest"

	"go_sample/app/infrastructure/middleware/echo/context"

	"github.com/labstack/echo/v4"
)

func AddChatsRoutingGroup(e *echo.Echo, ctrl *rest_controller.Controller) {
	chats := e.Group("chats")
	{
		// 対象ユーザーが所属するcannelを全て取得する
		chats.GET("", func(c echo.Context) error { return ctrl.Chat().FindAll(c.(*context.CustomContext)) })
		// 対象ユーザーが所属する単一cannelのメッセージを取得する（reload or 初期表示用、差分はwsで取得）
		chats.GET("/:id/messages", func(c echo.Context) error { return ctrl.Chat().FindMessages(c.(*context.CustomContext)) })

		// 更新系は即時対応なので、ws側で行う
		// chats.POST("", func(c echo.Context) error { return ctrl.Group().Create(c.(*context.CustomContext)) })
		// chats.PUT("/:id", func(c echo.Context) error { return ctrl.Group().Update(c.(*context.CustomContext)) })
		// chats.DELETE("/:id", func(c echo.Context) error { return ctrl.Group().Delete(c.(*context.CustomContext)) })
	}
}
