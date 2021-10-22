package mysql

import (
	"fmt"
	"go_sample/app/infrastructure/config"
	"log"

	"go_sample/app/infrastructure/middleware/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDb() *gorm.DB {
	config := config.LoadConfig()
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalln(connectionString + "database can't connect")
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.AutoMigrate(&model.UserRDBRecord{})
	db.AutoMigrate(&model.GroupRDBRecord{})

	return db
}
