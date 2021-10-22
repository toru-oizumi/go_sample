package input

import (
	"go_sample/app/domain/model"
)

type FindGroupByIdRequest struct {
	Id model.GroupId `validate:"required"`
}

type CreateGroupRequest struct {
	Name model.GroupName `validate:"required"`
}

type UpdateGroupRequest struct {
	Id   model.GroupId `validate:"required"`
	Name model.GroupName
}

type DeleteGroupByIdRequest struct {
	Id model.GroupId `validate:"required"`
}
