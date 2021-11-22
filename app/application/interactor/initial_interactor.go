package interactor

import (
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

// ALLの場合は最初に作成される？
// Groupの場合は、Group作成時に作成される
// DMの場合はメッセージ送信時に無ければ作成される

type InitialInteractor struct {
	Connection repository.Connection
}

func (i *InitialInteractor) DataBaseInitialize() error {
	// テーブルの作成
	i.Connection.Initialize().AutoMigrate()
	//	db.AutoMigrate(&model.UserRDBRecord{})
	// db.AutoMigrate(&model.GroupRDBRecord{})
	// db.AutoMigrate(&model.PlayRDBRecord{})

	// 未所属Groupの作成
	group := model.Group{
		Name: model.GroupName("free"),
	}

	i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_group, err := tx.Group().Store(group); err != nil {
				return nil, err
			} else {
				return created_group, nil
			}
		},
	)

	// ALL向けChatの作成
	return nil
}
