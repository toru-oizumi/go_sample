package input

import (
	"go_sample/app/domain/model"
)

type FindGroupByIDRequest struct {
	ID model.GroupID `param:"id" validate:"required"`
}

type FindGroupsRequest struct {
	NameLike string `query:"nameLike" validate:"required,omitempty,alphanum"`
}

type CreateGroupRequest struct {
	UserID model.UserID    `param:"userID" validate:"required"`
	Name   model.GroupName `json:"name" form:"name" validate:"required"`
}

type UpdateGroupRequest struct {
	ID   model.GroupID   `param:"id" validate:"required"`
	Name model.GroupName `json:"name" form:"name" validate:"required"`
}

type DeleteGroupRequest struct {
	ID model.GroupID `param:"id" validate:"required"`
}
