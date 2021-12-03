package gorm

import (
	mysql_service "go_sample/app/infrastructure/middleware/db/mysql"
	"go_sample/app/infrastructure/middleware/gorm/mysql"

	"go_sample/app/infrastructure/middleware/zap"
	batch_controller "go_sample/app/interface/controller/batch"
)

func Init() {
	db := mysql.NewDB()
	db_service := mysql_service.NewDBService()

	repository := NewRepository(db, db_service)
	connection, _ := repository.NewConnection()
	logger := zap.NewZapBatchLogger()

	controller := batch_controller.NewController(connection, logger)
	controller.Initial().Initialize()
}
