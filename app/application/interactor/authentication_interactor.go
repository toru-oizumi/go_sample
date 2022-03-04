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

type AuthenticationInteractor struct {
	Connection repository.Connection
	Presenter  presenter.AuthenticationPresenter
}

func (i *AuthenticationInteractor) SingIn(request input.SignInRequest) (*output.AuthenticationResponse, error) {
	if err := i.Connection.Account().Authenticate(request.Email, request.Password); err != nil {
		if errors.As(err, &util_error.ErrEntityNotExists{}) {
			return nil, util_error.NewErrAuthenticationFailed()
		} else {
			return nil, err
		}
	}

	account, err := i.Connection.Account().FindByEmail(request.Email)
	if err != nil {
		return nil, err
	}

	if user, err := i.Connection.User().FindByID(account.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildAuthenticationResponse(*user)
	}
}

func (i *AuthenticationInteractor) SignUp(request input.SignUpRequest) (*output.AuthenticationResponse, error) {
	user := model.User{
		Name: request.Name,
	}

	exists, err := i.Connection.User().ExistsByName(user.Name)
	if err != nil {
		return nil, err
	} else if exists {
		return nil, util_error.NewErrEntityAlreadyExists()
	}

	created_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			accountID, err := tx.Account().Store(
				model.Account{
					Email: request.Email,
				},
			)
			if err != nil {
				return nil, err
			}

			domain_service := service.NewDomainService(tx)

			user.ID = *accountID
			created_user_id, err := domain_service.User.Create(user)
			if err != nil {
				// Accountの場合だけ別途Deleteするのは、
				// Accountの実装がCognito依存なのが漏れて来ている…
				tx.Account().Delete(*accountID)
				return nil, err
			}

			created_user, err := tx.User().FindByID(*created_user_id)
			if err != nil {
				// Accountの場合だけ別途Deleteするのは、
				// Accountの実装がCognito依存なのが漏れて来ている…
				tx.Account().Delete(*accountID)
				return nil, err
			}

			return *created_user, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_user, _ := created_user.(model.User)
	return i.Presenter.BuildAuthenticationResponse(parsed_user)
}

func (i *AuthenticationInteractor) Activate(request input.ActivateRequest) (*output.AuthenticationResponse, error) {
	activated_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Account().Activate(request.Email, request.CurrentPassword, request.NewPassword); err != nil {
				return nil, err
			}
			account, err := tx.Account().FindByEmail(request.Email)
			if err != nil {
				return nil, err
			}
			activated_user, err := tx.User().FindByID(account.ID)
			if err != nil {
				return nil, err
			}
			return *activated_user, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_user, _ := activated_user.(model.User)
	return i.Presenter.BuildAuthenticationResponse(parsed_user)
}

func (i *AuthenticationInteractor) FindAccount(request input.FindAccountRequest) (*output.AuthenticationResponse, error) {
	if user, err := i.Connection.User().FindByID(request.UserID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildAuthenticationResponse(*user)
	}
}

func (i *AuthenticationInteractor) ChangePassword(request input.ChangePasswordRequest) (*output.AuthenticationResponse, error) {
	changed_user, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Account().ChangePassword(request.Email, request.CurrentPassword, request.NewPassword); err != nil {
				return nil, err
			}
			account, err := tx.Account().FindByEmail(request.Email)
			if err != nil {
				return nil, err
			}
			changed_user, err := tx.User().FindByID(account.ID)
			if err != nil {
				return nil, err
			}
			return *changed_user, nil
		},
	)

	if err != nil {
		return nil, err
	}

	parsed_user, _ := changed_user.(model.User)
	return i.Presenter.BuildAuthenticationResponse(parsed_user)
}

func (i *AuthenticationInteractor) SignOut(request input.SignOutRequest) error {
	return nil
}
