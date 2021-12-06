package echo

import (
	rest_router "go_sample/app/infrastructure/middleware/echo/router/rest"
	ws_router "go_sample/app/infrastructure/middleware/echo/router/ws"

	mysql_service "go_sample/app/infrastructure/middleware/db/mysql"
	"go_sample/app/infrastructure/middleware/echo/context"
	"go_sample/app/infrastructure/middleware/gorm"
	"go_sample/app/infrastructure/middleware/gorm/mysql"
	"go_sample/app/infrastructure/middleware/validator"
	"go_sample/app/infrastructure/middleware/zap"

	rest_controller "go_sample/app/interface/controller/rest"
	ws_controller "go_sample/app/interface/controller/ws"

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

	db := mysql.NewDB()
	db_service := mysql_service.NewDBService()

	repository := gorm.NewRepository(db, db_service)
	connection, _ := repository.NewConnection()

	logger := zap.NewZapApiResponseLogger()

	rest_ctrl := rest_controller.NewController(connection)
	ws_ctrl := ws_controller.NewWsHandler(connection)

	// 認証を行う
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("secret"),
	// }))
	// TODO: 最終的にはJWTの検証としたい
	// e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	// 	// return key == "valid-key", nil
	// 	return true, nil
	// }))

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

	rest_router.AddUsersRoutingGroup(e, rest_ctrl)
	rest_router.AddGroupsRoutingGroup(e, rest_ctrl)
	rest_router.AddFieldsRoutingGroup(e, rest_ctrl)
	rest_router.AddChatsRoutingGroup(e, rest_ctrl)

	ws_router.AddWsRoutingGroup(e, ws_ctrl)

	e.Logger.Fatal(e.Start(":18080"))
}
