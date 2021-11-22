package interactor

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type GroupInteractor struct {
	Connection repository.Connection
	Presenter  presenter.GroupPresenter
}

func (i *GroupInteractor) FindByID(request input.FindGroupByIDRequest) (*output.GroupResponse, error) {
	if group, err := i.Connection.Group().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildGroupResponse(*group)
	}
}

func (i *GroupInteractor) FindList(request input.FindGroupsRequest) ([]output.GroupResponse, error) {
	filter := repository.GroupFilter{
		NameLike: request.NameLike,
	}
	if users, err := i.Connection.Group().List(filter); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildGroupsResponse(users)
	}
}

func (i *GroupInteractor) FindAll() ([]output.GroupResponse, error) {
	if groups, err := i.Connection.Group().List(repository.GroupFilter{}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildGroupsResponse(groups)
	}
}

func (i *GroupInteractor) Create(request input.CreateGroupRequest) (*output.GroupResponse, error) {
	// TODO: Groupに所属していない場合しか作成できない
	group := model.Group{
		Name: request.Name,
	}

	if created_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			// Group作成
			created_group, err := tx.Group().Store(group)
			if err != nil {
				return nil, err
			}

			// Userの取得
			// ChatのMembersに設定するUserIDの存在確認も兼ねて、先に取得する
			user, err := tx.User().FindByID(request.UserID)
			if err != nil {
				return nil, err
			}

			// Group向けchatを作成
			chat := model.Chat{
				Name:    model.ChatName(request.Name),
				Members: []model.UserID{request.UserID},
			}
			if _, err := tx.Chat().Store(chat); err != nil {
				return nil, err
			}

			// userのGroupを書き換え
			user.Group = *created_group
			if _, err := tx.User().Update(*user); err != nil {
				return nil, err
			}

			return created_group, nil
		},
	); err != nil {
		return nil, err
	} else {
		parsed_group, _ := created_group.(model.Group)
		return i.Presenter.BuildGroupResponse(parsed_group)
	}
}

func (i *GroupInteractor) Update(request input.UpdateGroupRequest) (*output.GroupResponse, error) {
	group := model.Group{
		ID:   request.ID,
		Name: request.Name,
	}

	if updated_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_group, err := tx.Group().Update(group); err != nil {
				return nil, err
			} else {
				return updated_group, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_group, _ := updated_group.(model.Group)
		return i.Presenter.BuildGroupResponse(parsed_group)
	}
}

func (i *GroupInteractor) Delete(request input.DeleteGroupRequest) error {
	if _, err := i.Connection.Group().FindByID(request.ID); err != nil {
		return err
	}
	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Group().Delete(request.ID); err != nil {
				return nil, err
			} else {
				return nil, nil
			}

			// TODO: Group向けChatの削除
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}
