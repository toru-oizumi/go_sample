package input

import (
	"go_sample/app/domain/model"
)

type FindGroupByIDRequest struct {
	ID model.GroupID `param:"id" validate:"required"`
}

type CreateGroupRequest struct {
	Name model.GroupName `json:"name" form:"name" validate:"required"`
}

type UpdateGroupRequest struct {
	ID   model.GroupID   `param:"id" validate:"required"`
	Name model.GroupName `json:"name" form:"name" validate:"required"`
}

type DeleteGroupByIDRequest struct {
	ID model.GroupID `param:"id" validate:"required"`
}
