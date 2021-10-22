package echo

import (
	"fmt"

	"go_sample/app/infrastructure/middleware/echo/context"
	"go_sample/app/infrastructure/middleware/echo/router"
	"go_sample/app/infrastructure/middleware/gorm"
	"go_sample/app/infrastructure/middleware/gorm/mysql"
	"go_sample/app/infrastructure/middleware/zap"
	"go_sample/app/interface/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func handleWebSocket(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()

		// 初回のメッセージを送信
		err := websocket.Message.Send(ws, "Server: Hello, Client!")
		if err != nil {
			c.Logger().Error(err)
		}

		for {
			// Client からのメッセージを読み込む
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			// Client からのメッセージを元に返すメッセージを作成し送信する
			err := websocket.Message.Send(ws, fmt.Sprintf("Server: \"%s\" received!", msg))
			if err != nil {
				c.Logger().Error(err)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func Init() {
	e := echo.New()

	// アクセスログのようなリクエスト単位のログを出力する
	e.Use(middleware.Logger())
	// アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする
	e.Use(middleware.Recover())

	db := mysql.NewDb()
	repository := gorm.NewRepository(db)
	connection, _ := repository.NewConnection()

	logger := zap.NewZapLogger()

	ctrl := controller.NewController(connection, logger)

	// CustomContextを使用する
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{Context: c}
			return next(cc)
		}
	})

	router.AddUsersRoutingGroup(e, ctrl)
	router.AddGroupsRoutingGroup(e, ctrl)

	e.GET("/ws", handleWebSocket)

	e.Logger.Fatal(e.Start(":18080"))
}
