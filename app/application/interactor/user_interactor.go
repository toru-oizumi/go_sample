package interactor

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type UserInteractor struct {
	Connection repository.Connection
	Presenter  presenter.UserPresenter
}

func (i *UserInteractor) FindByID(request input.FindUserByIDRequest) (*output.FindUserByIDResponse, error) {
	if user, err := i.Connection.User().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindByIDResponse(user)
	}
}

func (i *UserInteractor) FindAll() (output.FindAllUsersResponse, error) {
	if users, err := i.Connection.User().List(repository.UserFilter{NameLike: ""}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindAllResponse(users)
	}
}

func (i *UserInteractor) Create(request input.CreateUserRequest) (*output.CreateUserResponse, error) {
	group, err := i.Connection.Group().FindByID(request.GroupID)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:  request.Name,
		Age:   request.Age,
		Group: *group,
	}

	if created_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_user, err := tx.User().Store(user); err != nil {
				return nil, err
			} else {
				return created_user, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_user, _ := created_user.(*model.User)
		return i.Presenter.BuildCreateResponse(parsed_user)
	}

}

func (i *UserInteractor) Update(request input.UpdateUserRequest) (*output.UpdateUserResponse, error) {
	group, err := i.Connection.Group().FindByID(request.GroupID)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:    request.ID,
		Name:  request.Name,
		Age:   request.Age,
		Group: *group,
	}

	if updated_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_user, err := tx.User().Update(user); err != nil {
				return nil, err
			} else {
				return updated_user, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_user, _ := updated_user.(*model.User)
		return i.Presenter.BuildUpdateResponse(parsed_user)
	}
}

func (i *UserInteractor) DeleteByID(request input.DeleteUserByIDRequest) error {
	if _, err := i.Connection.User().FindByID(request.ID); err != nil {
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.User().DeleteByID(request.ID); err != nil {
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
