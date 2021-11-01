package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type RoomPresenter interface {
	BuildFindByIDResponse(object *model.Room) (*output.FindRoomByIDResponse, error)
	BuildFindAllResponse(objects model.Rooms) (output.FindAllRoomsResponse, error)
	BuildCreateResponse(object *model.Room) (*output.CreateRoomResponse, error)
	BuildUpdateResponse(object *model.Room) (*output.UpdateRoomResponse, error)
}
