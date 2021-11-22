package input

import (
	"go_sample/app/domain/model"
)

type FindPlayByIDRequest struct {
	ID model.PlayID `param:"id" validate:"required"`
}

type CreatePlayRequest struct {
	Name          model.PlayName `json:"name" form:"name" validate:"required"`
	OwnerUserID   model.UserID   `json:"ownerUserID" form:"ownerUserID" validate:"required"`
	VisitorUserID model.UserID   `json:"visitorUserID" form:"visitorUserID" validate:"required"`
}

type UpdatePlayRequest struct {
	ID   model.PlayID   `param:"id" validate:"required"`
	Name model.PlayName `json:"name" form:"name" validate:"required"`
}

type DeletePlayRequest struct {
	ID model.PlayID `param:"id" validate:"required"`
}
