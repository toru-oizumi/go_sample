package echo

import (
	"go_sample/app/infrastructure/config"
	"go_sample/app/infrastructure/factory"
	rest_router "go_sample/app/infrastructure/middleware/echo/router/rest"
	ws_router "go_sample/app/infrastructure/middleware/echo/router/ws"

	mysql_service "go_sample/app/infrastructure/middleware/db/mysql"
	"go_sample/app/infrastructure/middleware/echo/context"
	"go_sample/app/infrastructure/middleware/gorm/mysql"
	"go_sample/app/infrastructure/middleware/validator"
	"go_sample/app/infrastructure/middleware/zap"

	rest_controller "go_sample/app/interface/controller/rest"
	ws_controller "go_sample/app/interface/controller/ws"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/wader/gormstore/v2"
)

func Init() {
	config := config.LoadConfig()

	e := echo.New()

	// アクセスログのようなリクエスト単位のログを出力する
	e.Use(middleware.Logger())
	// アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーする
	e.Use(middleware.Recover())

	e.Validator = validator.NewCustomValidator()

	db := mysql.NewDB(config)
	db_service := mysql_service.NewDBService()

	// gormを使用してセッションを管理する
	e.Use(session.Middleware(gormstore.New(db, []byte(config.SessionKey))))

	repository := factory.NewRepository(config, db, db_service)
	connection, _ := repository.NewConnection()

	logger := zap.NewZapApiResponseLogger()

	rest_ctrl := rest_controller.NewController(connection)
	ws_ctrl := ws_controller.NewWsController(connection)

	// CustomContextを使用する
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.CustomContext{
				Context:       c,
				RestApiLogger: logger,
			}
			return next(cc)
		}
	})

	rest_router.AddAuthenticationRoutingGroup(e, rest_ctrl)
	rest_router.AddUsersRoutingGroup(e, rest_ctrl)
	rest_router.AddGroupsRoutingGroup(e, rest_ctrl)
	rest_router.AddFieldsRoutingGroup(e, rest_ctrl)
	rest_router.AddChatsRoutingGroup(e, rest_ctrl)
	rest_router.AddDirectMessagesRoutingGroup(e, rest_ctrl)

	ws_router.AddWsRoutingGroup(e, ws_ctrl)

	e.Logger.Fatal(e.Start(":" + config.RunServerPort))
}
