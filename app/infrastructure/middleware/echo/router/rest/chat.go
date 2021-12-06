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
		chats.GET("/:chatID/messages", func(c echo.Context) error { return ctrl.Chat().FindMessages(c.(*context.CustomContext)) })
	}
}
