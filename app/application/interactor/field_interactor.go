package interactor

import (
	"errors"
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	util_error "go_sample/app/utility/error"
)

type FieldInteractor struct {
	Connection repository.Connection
	Presenter  presenter.FieldPresenter
}

func (i *FieldInteractor) FindByID(request input.FindFieldByIDRequest) (*output.FieldResponse, error) {
	if field, err := i.Connection.Field().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFieldResponse(*field)
	}
}

func (i *FieldInteractor) FindAll() ([]output.FieldResponse, error) {
	if fields, err := i.Connection.Field().List(repository.FieldFilter{FieldID: ""}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFieldsResponse(fields)
	}
}

func (i *FieldInteractor) Create(request input.CreateFieldRequest) (*output.FieldResponse, error) {
	owner_user, err := i.Connection.User().FindByID(request.OwnerUserID)
	if err != nil {
		return nil, err
	}

	visitor_user, err := i.Connection.User().FindByID(request.VisitorUserID)
	if err != nil {
		return nil, err
	}

	field := model.Field{
		Name:          request.Name,
		OwnerUserID:   owner_user.ID,
		VisitorUserID: visitor_user.ID,
	}

	if created_field, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_field, err := tx.Field().Store(field); err != nil {
				return nil, err
			} else {
				return created_field, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_field, _ := created_field.(model.Field)
		return i.Presenter.BuildFieldResponse(parsed_field)
	}

}

func (i *FieldInteractor) Update(request input.UpdateFieldRequest) (*output.FieldResponse, error) {
	field := model.Field{
		ID:   request.ID,
		Name: request.Name,
	}

	if updated_field, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_field, err := tx.Field().Update(field); err != nil {
				return nil, err
			} else {
				return updated_field, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_field, _ := updated_field.(model.Field)
		return i.Presenter.BuildFieldResponse(parsed_field)
	}
}

func (i *FieldInteractor) Delete(request input.DeleteFieldRequest) error {
	if _, err := i.Connection.Field().FindByID(request.ID); err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrRecordNotFound{}) {
			return nil
		}
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			err := tx.Field().Delete(request.ID)
			return nil, err
		},
	); err != nil {
		return err
	} else {
		return nil
	}
}
