package repository_impl

import (
	"go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"

	"gorm.io/gorm"
)

type Initialize struct {
	DB      *gorm.DB
	Service db.DBService
}

func (i *Initialize) AutoMigrate() error {
	if err := i.DB.AutoMigrate(&model.UserRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.GroupRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.FieldRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.ChatRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.ChatMessageRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.UserChatRDBRecord{}); err != nil {
		return err
	}
	if err := i.DB.AutoMigrate(&model.DirectMessageRDBRecord{}); err != nil {
		return err
	}
	return nil
}
