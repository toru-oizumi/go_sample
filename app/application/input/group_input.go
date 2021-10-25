package input

import (
	"go_sample/app/domain/model"

	"gopkg.in/go-playground/validator.v9"
)

type FindGroupByIdRequest struct {
	Id model.GroupId `validate:"required"`
}

func (r *FindGroupByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type CreateGroupRequest struct {
	Name model.GroupName `validate:"required"`
}

func (r *CreateGroupRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UpdateGroupRequest struct {
	Id   model.GroupId   `validate:"required"`
	Name model.GroupName `validate:"required"`
}

func (r *UpdateGroupRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type DeleteGroupByIdRequest struct {
	Id model.GroupId `validate:"required"`
}

func (r *DeleteGroupByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
