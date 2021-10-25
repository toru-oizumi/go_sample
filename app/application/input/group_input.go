package input

import (
	"go_sample/app/domain/model"
)

type FindGroupByIdRequest struct {
	Id model.GroupId `param:"id" validate:"required"`
}

type CreateGroupRequest struct {
	Name model.GroupName `json:"name" form:"name" validate:"required"`
}

type UpdateGroupRequest struct {
	Id   model.GroupId   `param:"id" validate:"required"`
	Name model.GroupName `json:"name" form:"name" validate:"required"`
}

type DeleteGroupByIdRequest struct {
	Id model.GroupId `param:"id" validate:"required"`
}
