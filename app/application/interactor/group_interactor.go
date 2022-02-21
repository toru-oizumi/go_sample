package interactor

import (
	"errors"
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	"go_sample/app/domain/service"
	util_error "go_sample/app/utility/error"
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
	group := model.Group{
		Name:            request.Name,
		NumberOfMembers: 0,
	}

	// Userの取得
	// ChatのMembersに設定するUserの存在確認も兼ねて、先に取得する
	user, err := i.Connection.User().FindByID(request.UserID)
	if err != nil {
		return nil, err
	}

	// グループに未所属(FreeGroupNameのグループに所属)の場合のみ、グループを作成できる
	if user.Group.Name != model.FreeGroupName {
		return nil, util_error.NewErrBadRequest("only users who are not yet members of a group can create a group")
	}

	created_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)

			// Group作成
			created_group_id, err := domain_service.Group.Create(group)
			if err != nil {
				return nil, err
			}

			created_group, err := tx.Group().FindByID(*created_group_id)
			if err != nil {
				return nil, err
			}

			if err := domain_service.User.JoinGroup(*user, *created_group); err != nil {
				return nil, err
			}

			return *created_group, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_group, _ := created_group.(model.Group)
	return i.Presenter.BuildGroupResponse(parsed_group)
}

func (i *GroupInteractor) Update(request input.UpdateGroupRequest) (*output.GroupResponse, error) {
	group := model.Group{
		ID:   request.ID,
		Name: request.Name,
	}

	updated_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)

			// Group更新
			updated_group_id, err := domain_service.Group.Update(group)
			if err != nil {
				return nil, err
			}

			updated_group, err := tx.Group().FindByID(*updated_group_id)
			if err != nil {
				return nil, err
			}

			return *updated_group, nil
		},
	)
	if err != nil {
		return nil, err
	}

	parsed_group, _ := updated_group.(model.Group)
	return i.Presenter.BuildGroupResponse(parsed_group)
}

func (i *GroupInteractor) Delete(request input.DeleteGroupRequest) error {
	if _, err := i.Connection.Group().FindByID(request.ID); err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrEntityNotExists{}) {
			return nil
		}
		return err
	}

	_, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)

			err := domain_service.Group.Delete(request.ID)
			return nil, err
		},
	)

	if err != nil {
		return err
	} else {
		return nil
	}
}
