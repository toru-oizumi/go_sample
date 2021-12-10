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

type UserInteractor struct {
	Connection repository.Connection
	Presenter  presenter.UserPresenter
}

func (i *UserInteractor) FindByID(request input.FindUserByIDRequest) (*output.UserResponse, error) {
	if user, err := i.Connection.User().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildUserResponse(*user)
	}
}

func (i *UserInteractor) FindList(request input.FindUsersRequest) ([]output.UserResponse, error) {
	filter := repository.UserFilter{
		GroupID:  request.GroupID,
		NameLike: request.NameLike,
	}
	if users, err := i.Connection.User().List(filter); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildUsersResponse(users)
	}
}

func (i *UserInteractor) FindAll() ([]output.UserResponse, error) {
	if users, err := i.Connection.User().List(repository.UserFilter{}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildUsersResponse(users)
	}
}

func (i *UserInteractor) Create(request input.CreateUserRequest) (*output.UserResponse, error) {
	user := model.User{
		Name: request.Name,
	}

	created_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)

			created_user_id, err := domain_service.User.Create(user)
			if err != nil {
				return nil, err
			}

			created_user, err := tx.User().FindByID(*created_user_id)
			if err != nil {
				return nil, err
			}

			return *created_user, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_user, _ := created_user.(model.User)
	return i.Presenter.BuildUserResponse(parsed_user)
}

func (i *UserInteractor) Update(request input.UpdateUserRequest) (*output.UserResponse, error) {
	after_group, err := i.Connection.Group().FindByID(request.GroupID)
	if err != nil {
		return nil, err
	}

	current_user, err := i.Connection.User().FindByID(request.ID)
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:    request.ID,
		Name:  request.Name,
		Group: current_user.Group, // Groupは変更がある場合はJoinGroupで更新するので、ここではcurrent_user.Groupを設定
	}

	updated_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			updated_user_id, err := tx.User().Update(user)
			if err != nil {
				return nil, err
			}

			// Groupを変更した場合
			if current_user.Group.ID != request.GroupID {
				domain_service := service.NewDomainService(tx)
				if err := domain_service.User.JoinGroup(user, *after_group); err != nil {
					return nil, err
				}
			}

			updated_user, err := tx.User().FindByID(*updated_user_id)
			if err != nil {
				return nil, err
			}

			return *updated_user, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_user, _ := updated_user.(model.User)
	return i.Presenter.BuildUserResponse(parsed_user)
}

func (i *UserInteractor) Delete(request input.DeleteUserRequest) error {
	user, err := i.Connection.User().FindByID(request.ID)
	if err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrRecordNotFound{}) {
			return nil
		}
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			domain_service := service.NewDomainService(tx)
			err := domain_service.User.Delete(*user)
			return nil, err
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}
