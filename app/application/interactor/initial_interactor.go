package interactor

import (
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	"go_sample/app/domain/service"
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

	// 未所属Groupの作成
	group := model.Group{
		Name: model.FreeGroupName,
	}

	// 全体用Chatの作成
	chat := model.Chat{
		Name: model.AllChatName,
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)
			if _, err := domain_service.Group.Create(group); err != nil {
				return nil, err
			}
			if _, err := tx.Chat().Store(chat); err != nil {
				return nil, err
			}
			return nil, nil
		},
	); err != nil {
		return err
	}
	return nil
}
