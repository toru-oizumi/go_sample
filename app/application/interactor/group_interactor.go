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

func (i *GroupInteractor) FindByID(request input.FindGroupByIDRequest) (*output.FindGroupByIDResponse, error) {
	if group, err := i.Connection.Group().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindByIDResponse(group)
	}
}

func (i *GroupInteractor) FindAll() (output.FindAllGroupsResponse, error) {
	if groups, err := i.Connection.Group().List(repository.GroupFilter{UserID: model.UserID("1")}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindAllResponse(groups)
	}
}

func (i *GroupInteractor) Create(request input.CreateGroupRequest) (*output.CreateGroupResponse, error) {
	group := model.Group{
		Name: request.Name,
	}

	if created_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_group, err := tx.Group().Store(group); err != nil {
				return nil, err
			} else {
				return created_group, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_group, _ := created_group.(*model.Group)
		return i.Presenter.BuildCreateResponse(parsed_group)
	}
}

func (i *GroupInteractor) Update(request input.UpdateGroupRequest) (*output.UpdateGroupResponse, error) {
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
		parsed_group, _ := updated_group.(*model.Group)
		return i.Presenter.BuildUpdateResponse(parsed_group)
	}
}

func (i *GroupInteractor) DeleteByID(request input.DeleteGroupByIDRequest) error {
	if _, err := i.Connection.Group().FindByID(request.ID); err != nil {
		return err
	}
	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Group().DeleteByID(request.ID); err != nil {
				return nil, err
			} else {
				return nil, nil
			}
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}
