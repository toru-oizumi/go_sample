package input

import (
	"go_sample/app/domain/model"

	"gopkg.in/go-playground/validator.v9"
)

type FindRoomByIdRequest struct {
	Id model.RoomId `validate:"required"`
}

func (r *FindRoomByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type CreateRoomRequest struct {
	Name          model.RoomName `validate:"required"`
	OwnerUserId   model.UserId   `validate:"required"`
	VisitorUserId model.UserId   `validate:"required"`
}

func (r *CreateRoomRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UpdateRoomRequest struct {
	Id   model.RoomId   `validate:"required"`
	Name model.RoomName `validate:"required"`
}

func (r *UpdateRoomRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type DeleteRoomByIdRequest struct {
	Id model.RoomId `validate:"required"`
}

func (r *DeleteRoomByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
