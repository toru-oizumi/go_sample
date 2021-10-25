package echo

import (
	ws_router "go_sample/app/infrastructure/middleware/echo/router/ws"

	"go_sample/app/infrastructure/middleware/echo/context"
	"go_sample/app/infrastructure/middleware/echo/router"
	"go_sample/app/infrastructure/middleware/gorm"
	"go_sample/app/infrastructure/middleware/gorm/mysql"
	"go_sample/app/infrastructure/middleware/validator"
	"go_sample/app/infrastructure/middleware/zap"

	rest_controller "go_sample/app/interface/controller"
	ws_handler "go_sample/app/interface/controller/ws"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	e := echo.New()

	// アクセスログのようなリクエスト単位のログを出力する
	e.Use(middleware.Logger())
	// アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする
	e.Use(middleware.Recover())

	e.Validator = validator.NewCustomValidator()

	db := mysql.NewDb()
	repository := gorm.NewRepository(db)
	connection, _ := repository.NewConnection()

	logger := zap.NewZapLogger()

	controller := rest_controller.NewController(connection, logger)
	handler := ws_handler.NewWsHandler(connection, logger)

	// CustomContextを使用する
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{Context: c}
			return next(cc)
		}
	})

	router.AddUsersRoutingGroup(e, controller)
	router.AddGroupsRoutingGroup(e, controller)
	router.AddRoomsRoutingGroup(e, controller)

	ws_router.AddWsRoomsRoutingGroup(e, handler)

	e.Logger.Fatal(e.Start(":18080"))
}
