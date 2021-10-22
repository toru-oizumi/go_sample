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

func (i *GroupInteractor) FindById(request input.FindGroupByIdRequest) (*output.FindGroupByIdResponse, error) {
	if group, err := i.Connection.Group().FindById(request.Id); err != nil {
		return nil, err
	} else {
		return i.Presenter.CreateFindGroupByIdResponse(group)
	}
}

func (i *GroupInteractor) FindAll() (output.FindAllGroupsResponse, error) {
	if groups, err := i.Connection.Group().List(repository.GroupFilter{UserID: model.UserId(1)}); err != nil {
		return nil, err
	} else {
		return i.Presenter.CreatFindAllGroupsResponse(groups)
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
		return i.Presenter.CreateCreateGroupResponse(parsed_group)
	}
}

func (i *GroupInteractor) Update(request input.UpdateGroupRequest) (*output.UpdateGroupResponse, error) {
	group, err := i.Connection.Group().FindById(request.Id)
	if err != nil {
		return nil, err
	}

	group.Name = request.Name

	if updated_group, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_group, err := tx.Group().Update(*group); err != nil {
				return nil, err
			} else {
				return updated_group, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_group, _ := updated_group.(*model.Group)
		return i.Presenter.CreateUpdateGroupResponse(parsed_group)
	}
}

func (i *GroupInteractor) DeleteById(request input.DeleteGroupByIdRequest) error {
	if _, err := i.Connection.Group().FindById(request.Id); err != nil {
		return err
	}
	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Group().DeleteById(request.Id); err != nil {
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