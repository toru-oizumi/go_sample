package gorm

import (
	"go_sample/app/infrastructure/config"
	"go_sample/app/infrastructure/factory"
	mysql_service "go_sample/app/infrastructure/middleware/db/mysql"
	"go_sample/app/infrastructure/middleware/gorm/mysql"

	"go_sample/app/infrastructure/middleware/zap"
	batch_controller "go_sample/app/interface/controller/batch"
)

func Init() {
	config := config.LoadConfig()
	db := mysql.NewDB(config)
	db_service := mysql_service.NewDBService()

	repository := factory.NewRepository(config, db, db_service)
	connection, _ := repository.NewConnection()
	logger := zap.NewZapBatchLogger()

	controller := batch_controller.NewController(connection, logger)
	controller.Initial().Initialize()
}
