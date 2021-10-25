package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type RoomUsecase interface {
	FindById(request input.FindRoomByIdRequest) (*output.FindRoomByIdResponse, error)
	FindAll() (output.FindAllRoomsResponse, error)
	Create(request input.CreateRoomRequest) (*output.CreateRoomResponse, error)
	Update(request input.UpdateRoomRequest) (*output.UpdateRoomResponse, error)
	DeleteById(request input.DeleteRoomByIdRequest) error
}
