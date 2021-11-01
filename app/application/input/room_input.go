package input

import (
	"go_sample/app/domain/model"
)

type FindRoomByIDRequest struct {
	ID model.RoomID `param:"id" validate:"required"`
}

type CreateRoomRequest struct {
	Name          model.RoomName `json:"name" form:"name" validate:"required"`
	OwnerUserID   model.UserID   `json:"ownerUserID" form:"ownerUserID" validate:"required"`
	VisitorUserID model.UserID   `json:"visitorUserID" form:"visitorUserID" validate:"required"`
}

type UpdateRoomRequest struct {
	ID   model.RoomID   `param:"id" validate:"required"`
	Name model.RoomName `json:"name" form:"name" validate:"required"`
}

type DeleteRoomByIDRequest struct {
	ID model.RoomID `param:"id" validate:"required"`
}
