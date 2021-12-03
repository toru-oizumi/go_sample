package input

import (
	"go_sample/app/domain/model"
)

type FindFieldByIDRequest struct {
	ID model.FieldID `param:"id" validate:"required"`
}

type CreateFieldRequest struct {
	Name          model.FieldName `json:"name" form:"name" validate:"required"`
	OwnerUserID   model.UserID    `json:"ownerUserID" form:"ownerUserID" validate:"required"`
	VisitorUserID model.UserID    `json:"visitorUserID" form:"visitorUserID" validate:"required"`
}

type UpdateFieldRequest struct {
	ID   model.FieldID   `param:"id" validate:"required"`
	Name model.FieldName `json:"name" form:"name" validate:"required"`
}

type DeleteFieldRequest struct {
	ID model.FieldID `param:"id" validate:"required"`
}
