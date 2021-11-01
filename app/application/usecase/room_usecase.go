package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type RoomUsecase interface {
	FindByID(request input.FindRoomByIDRequest) (*output.FindRoomByIDResponse, error)
	FindAll() (output.FindAllRoomsResponse, error)
	Create(request input.CreateRoomRequest) (*output.CreateRoomResponse, error)
	Update(request input.UpdateRoomRequest) (*output.UpdateRoomResponse, error)
	DeleteByID(request input.DeleteRoomByIDRequest) error
}
