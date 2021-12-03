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

type PlayInteractor struct {
	Connection repository.Connection
	Presenter  presenter.PlayPresenter
}

func (i *PlayInteractor) FindByID(request input.FindPlayByIDRequest) (*output.PlayResponse, error) {
	if room, err := i.Connection.Play().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildPlayResponse(*room)
	}
}

func (i *PlayInteractor) FindAll() ([]output.PlayResponse, error) {
	if rooms, err := i.Connection.Play().List(repository.PlayFilter{PlayID: ""}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildPlaysResponse(rooms)
	}
}

func (i *PlayInteractor) Create(request input.CreatePlayRequest) (*output.PlayResponse, error) {
	owner_user, err := i.Connection.User().FindByID(request.OwnerUserID)
	if err != nil {
		return nil, err
	}

	visitor_user, err := i.Connection.User().FindByID(request.VisitorUserID)
	if err != nil {
		return nil, err
	}

	room := model.Play{
		Name:          request.Name,
		OwnerUserID:   owner_user.ID,
		VisitorUserID: visitor_user.ID,
	}

	if created_room, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_room, err := tx.Play().Store(room); err != nil {
				return nil, err
			} else {
				return created_room, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_room, _ := created_room.(model.Play)
		return i.Presenter.BuildPlayResponse(parsed_room)
	}

}

func (i *PlayInteractor) Update(request input.UpdatePlayRequest) (*output.PlayResponse, error) {
	room := model.Play{
		ID:   request.ID,
		Name: request.Name,
	}

	if updated_room, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_room, err := tx.Play().Update(room); err != nil {
				return nil, err
			} else {
				return updated_room, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_room, _ := updated_room.(model.Play)
		return i.Presenter.BuildPlayResponse(parsed_room)
	}
}

func (i *PlayInteractor) Delete(request input.DeletePlayRequest) error {
	if _, err := i.Connection.Play().FindByID(request.ID); err != nil {
		// 冪等性を重視して、削除の場合はrecord not foundエラーにしない
		if errors.As(err, &util_error.ErrRecordNotFound{}) {
			return nil
		}
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Play().Delete(request.ID); err != nil {
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
