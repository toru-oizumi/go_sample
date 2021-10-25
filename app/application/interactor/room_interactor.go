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

func (i *RoomInteractor) FindById(request input.FindRoomByIdRequest) (*output.FindRoomByIdResponse, error) {
	if room, err := i.Connection.Room().FindById(request.Id); err != nil {
		return nil, err
	} else {
		return i.Presenter.BuildFindByIdResponse(room)
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
	owner_user, err := i.Connection.User().FindById(request.OwnerUserId)
	if err != nil {
		return nil, err
	}

	visitor_user, err := i.Connection.User().FindById(request.VisitorUserId)
	if err != nil {
		return nil, err
	}

	room := model.Room{
		Name:          request.Name,
		OwnerUserId:   owner_user.Id,
		VisitorUserId: visitor_user.Id,
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
		Id:   request.Id,
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

func (i *RoomInteractor) DeleteById(request input.DeleteRoomByIdRequest) error {
	if _, err := i.Connection.Room().FindById(request.Id); err != nil {
		return err
	}

	if _, err := i.Connection.RunTransaction(
		func(tx repository.Transaction) (interface{}, error) {
			if err := tx.Room().DeleteById(request.Id); err != nil {
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
