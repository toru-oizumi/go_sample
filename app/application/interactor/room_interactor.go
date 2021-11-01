package interactor

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
	"go_sample/app/application/presenter"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
)

type RoomInteractor struct {
	Connection repository.Connection
	Presenter  presenter.RoomPresenter
}

func (i *RoomInteractor) FindByID(request input.FindRoomByIDRequest) (*output.FindRoomByIDResponse, error) {
	if room, err := i.Connection.Room().FindByID(request.ID); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindByIDResponse(room)
	}
}

func (i *RoomInteractor) FindAll() (output.FindAllRoomsResponse, error) {
	if rooms, err := i.Connection.Room().List(repository.RoomFilter{RoomID: ""}); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindAllResponse(rooms)
	}
}

func (i *RoomInteractor) Create(request input.CreateRoomRequest) (*output.CreateRoomResponse, error) {
	owner_user, err := i.Connection.User().FindByID(request.OwnerUserID)
	if err != nil {
		return nil, err
	}

	visitor_user, err := i.Connection.User().FindByID(request.VisitorUserID)
	if err != nil {
		return nil, err
	}

	room := model.Room{
		Name:          request.Name,
		OwnerUserID:   owner_user.ID,
		VisitorUserID: visitor_user.ID,
	}

	if created_room, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if created_room, err := tx.Room().Store(room); err != nil {
				return nil, err
			} else {
				return created_room, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_room, _ := created_room.(*model.Room)
		return i.Presenter.BuildCreateResponse(parsed_room)
	}

}

func (i *RoomInteractor) Update(request input.UpdateRoomRequest) (*output.UpdateRoomResponse, error) {
	room := model.Room{
		ID:   request.ID,
		Name: request.Name,
	}

	if updated_room, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if updated_room, err := tx.Room().Update(room); err != nil {
				return nil, err
			} else {
				return updated_room, nil
			}
		},
	); err != nil {
		return nil, err
	} else {
		parsed_room, _ := updated_room.(*model.Room)
		return i.Presenter.BuildUpdateResponse(parsed_room)
	}
}

func (i *RoomInteractor) DeleteByID(request input.DeleteRoomByIDRequest) error {
	if _, err := i.Connection.Room().FindByID(request.ID); err != nil {
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Room().DeleteByID(request.ID); err != nil {
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
