package input

import (
	"go_sample/app/domain/model"
)

type FindRoomByIdRequest struct {
	Id model.RoomId `param:"id" validate:"required"`
}

type CreateRoomRequest struct {
	Name          model.RoomName `json:"name" form:"name" validate:"required"`
	OwnerUserId   model.UserId   `json:"ownerUserId" form:"ownerUserId" validate:"required"`
	VisitorUserId model.UserId   `json:"visitorUserId" form:"visitorUserId" validate:"required"`
}

type UpdateRoomRequest struct {
	Id   model.RoomId   `param:"id" validate:"required"`
	Name model.RoomName `json:"name" form:"name" validate:"required"`
}

type DeleteRoomByIdRequest struct {
	Id model.RoomId `param:"id" validate:"required"`
}
